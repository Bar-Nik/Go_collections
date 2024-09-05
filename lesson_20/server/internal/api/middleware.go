package api

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func Logging(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var metod string
			switch r.Method {
			case "GET":
				metod = "получаем"
			case "DELETE":
				metod = "удаляем"
			case "POST":
				metod = "создаем"
			case "PUT":
				metod = "обновляем"
			}
			// body, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	handleError(w, http.StatusInternalServerError, err)
			// 	return
			// }
			logger.Info("Request",
				slog.String("msg", metod),
			)
			// logger.Info("Request",
			// 	slog.String("тело", string(body)),
			// )
			next.ServeHTTP(w, r)
			logger.Info("Finished")
		})
	}
}

// При логировании body с помощью io.ReadAll(r.Body) тело считывается в функции Logging,
// но при повторном считывании body с помощью io.ReadAll(r.Body) (в обработчиках) данных не будет.
