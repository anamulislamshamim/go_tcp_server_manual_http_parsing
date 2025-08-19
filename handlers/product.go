package handlers

import (
	"crud_2/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// HandleProducts routes all HTTP requests related to /products
func HandleProducts(method, path, body string) string {
	switch method {
	case "GET":
		// GET /products -> return all products
		if path == "/products" {
			response, _ := json.Marshal(models.Products)
			// fmt.Println("handler res: ", response)
			return HttpResponse(200, string(response))
		}

		return HttpResponse(400, `{"error":"I think it was Bad reqquest"}`)

	case "POST":
		// POST /products -> create new product
		if path == "/products" {
			var product models.Product
			// Decode JSON body directly into Product struct
			err := json.Unmarshal([]byte(strings.TrimSpace(body)), &product)
			if err != nil {
				fmt.Println("❌ JSON parse error:", err) // debug log
				return HttpResponse(400, `{"error":"invalid JSON"}`)
			}
			product.ID = models.NextId
			models.NextId++
			models.Products = append(models.Products, product)

			response, _ := json.Marshal(product)
			return HttpResponse(201, string(response))
		}

	case "PUT":
		// PUT /products/{id} -> update product
		parts := strings.Split(path, "/")
		if len(parts) == 3 {
			id, _ := strconv.Atoi(parts[2])
			var updated models.Product
			err := json.Unmarshal([]byte(body), &updated)
			if err != nil {
				return HttpResponse(400, `{"error":"invalid JSON"}`)
			}

			for i := range models.Products {
				if models.Products[i].ID == id {
					// update fields
					models.Products[i].Name = updated.Name
					models.Products[i].Price = updated.Price

					response, _ := json.Marshal(models.Products[i])
					return HttpResponse(200, string(response))
				}
			}
			return HttpResponse(404, `{"error":"product not found"}`)
		}

	case "DELETE":
		// DELETE /products/{id}
		parts := strings.Split(path, "/")
		if len(parts) == 3 {
			id, _ := strconv.Atoi(parts[2])
			for i := range models.Products {
				if models.Products[i].ID == id {
					// Remove element from slice
					models.Products = append(models.Products[:i], models.Products[i+1:]...)
					return HttpResponse(204, `{"message":"deleted"}`)
				}
			}
			return HttpResponse(404, `{"error":"product not found"}`)
		}
	}

	// Default: not found
	return HttpResponse(404, `{"error":"not found"}`)
}

// HttpResponse builds a raw HTTP response string
func HttpResponse(status int, body string) string {
	statusText := map[int]string{
		200: "OK",
		204: "No Content",
		201: "Created",
		400: "Bad Request",
		404: "Not Found",
	}[status]

	return fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s",
		status, statusText, len(body), body,
	)
}
