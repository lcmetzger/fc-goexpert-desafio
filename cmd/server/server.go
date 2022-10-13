package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Abre a conexão com o SQLite3
	db, err := sql.Open("sqlite3", "./cambio.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	err = createTableIfNotExists(db)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()
		cotacao, err := buscaCotacao(ctx)
		if err != nil {
			log.Println("Erro ao buscar cotacao", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*10)
		err = persisteCotacao(ctx, db, cotacao)
		defer cancel()
		if err != nil {
			log.Println("Erro ao persistir cotacao", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao buscar dados da API"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("bid: " + strconv.FormatFloat(float64(cotacao), 'f', 2, 32))
	})

	http.ListenAndServe(":8080", nil)

}

func buscaCotacao(ctx context.Context) (float64, error) {
	log.Println("Buscando dados da API")

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	cotacao, _ := strconv.ParseFloat(data["USDBRL"].(map[string]interface{})["bid"].(string), 64)
	return cotacao, nil
}

func persisteCotacao(ctx context.Context, db *sql.DB, cotacao float64) error {
	log.Println("Persistindo cotacao", time.Now(), cotacao)
	sql := `
			INSERT INTO cambio (data, cotacao)
			VALUES (?, ?)
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, time.Now(), cotacao)
	return err
}

func createTableIfNotExists(db *sql.DB) error {
	log.Println("Criando tabela se não existir")
	sql := `
			CREATE TABLE IF NOT EXISTS cambio (
				data  DATE,
				cotacao REAL
			);
	`
	_, err := db.Exec(sql)
	return err
}
