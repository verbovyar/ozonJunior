package validation

import "net/http"

func ReadValidate(key string, writer http.ResponseWriter) bool {
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)

		return false
	}

	return true
}

func CreateValidate(key string, value string, writer http.ResponseWriter) bool {
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)

		return false
	}

	if value == "" {
		writer.WriteHeader(http.StatusBadRequest)

		return false
	}

	return true
}

func UpdateValidation(key string, value string, writer http.ResponseWriter) bool {
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)

		return false
	}

	if value == "" {
		writer.WriteHeader(http.StatusNotFound)

		return false
	}

	return true
}

func DeleteValidate(key string, writer http.ResponseWriter) bool {
	if key == "" {
		writer.WriteHeader(http.StatusBadRequest)

		return false
	}

	return true
}
