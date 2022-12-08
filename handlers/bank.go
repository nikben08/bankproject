package handlers

import (
	"bankproject/models"
	"bankproject/seeds"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) CreateBank(w http.ResponseWriter, r *http.Request) {
	type CreateBankRequest struct {
		Name string `json:"name"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request CreateBankRequest
	json.Unmarshal(body, &request)

	bank := Bank{Name: request.Name}
	if result := h.DB.Create(&bank); result.Error != nil {
		json.NewEncoder(w).Encode("Couldnt create bank")
	}

	type Data struct {
		Id   int
		Name string
	}

	type Response struct {
		Data    Data
		Code    int
		Message string
	}

	response := Response{
		Data:    Data{Id: bank.ID, Name: bank.Name},
		Code:    200,
		Message: "Bank successfully created"}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h handler) AddInterest(w http.ResponseWriter, r *http.Request) {
	type AddInterestRequest struct {
		BankID     uint    `json:"bank_id"`
		Interest   float32 `json:"interest"`
		TimeOption int     `json:"time_option"`
		CreditType int     `json:"credit_type"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request AddInterestRequest
	json.Unmarshal(body, &request)

	var interest = models.Interest{
		BankID:     request.BankID,
		Interest:   request.Interest,
		TimeOption: request.TimeOption,
		CreditType: request.CreditType,
	}

	if result := h.DB.Create(&interest); result.Error != nil {
		json.NewEncoder(w).Encode("Couldnt create interest")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Interest successfully added")
	}
}

func (h handler) DeleteInterest(w http.ResponseWriter, r *http.Request) {
	type DeleteInterestRequest struct {
		InterestId int  `json:"id"`
		BankId     uint `json:"bank_id"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request DeleteInterestRequest
	json.Unmarshal(body, &request)

	if result := h.DB.Where("id = ? AND bank_id = ?", request.InterestId, request.BankId).Delete(&models.Interest{}); result.Error != nil {
		json.NewEncoder(w).Encode("Couldnt delete interest")
	}

	type Response struct {
		Code    string
		Message string
	}
	response := Response{Code: "200", Message: "Interest successfully deleted"}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h handler) RelevantInterests(w http.ResponseWriter, r *http.Request) {
	time_option, _ := strconv.Atoi(r.URL.Query().Get("time_option"))
	credit_type, _ := strconv.Atoi(r.URL.Query().Get("credit_type"))
	var interests []models.Interest
	if result := h.DB.Where("time_option = ? AND credit_type = ?", time_option, credit_type).Order("interest asc").Find(&interests); result.Error != nil {
		json.NewEncoder(w).Encode("Couldnt get interests")
	}

	type Row struct {
		BankId     uint
		Id         int
		Interest   float32
		TimeOption string
		CreditType string
	}

	var response []Row
	for i := 0; i < len(interests); i++ {
		var row Row
		row.BankId = interests[i].BankID
		row.Id = interests[i].ID
		row.Interest = interests[i].Interest
		row.CreditType = seeds.Credits()[credit_type]
		row.TimeOption = seeds.TimeOptions()[time_option]
		response = append(response, row)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h handler) GetBankById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankId := params["id"]
	id, _ := strconv.Atoi(bankId)

	var bank Bank
	if result := h.DB.First(&bank, id); result.Error != nil {
		json.NewEncoder(w).Encode("couldn't find bank")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bank)
	}
}

func (h handler) DeleteBankById(w http.ResponseWriter, r *http.Request) {
	type DeleteBankRequest struct {
		Id int `json:"id"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request DeleteBankRequest
	json.Unmarshal(body, &request)

	var bank Bank
	if result := h.DB.Unscoped().Delete(&bank, request.Id); result.Error != nil {
		json.NewEncoder(w).Encode("couldn't delete bank")
	} else {
		type Response struct {
			Code    string
			Message string
		}
		response := Response{Code: "200", Message: "Bank successfully deleted"}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (h handler) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	var banks []models.Bank
	if result := h.DB.Find(&banks); result.Error != nil {
		json.NewEncoder(w).Encode("While getting all banks error appeared")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(banks)
}
