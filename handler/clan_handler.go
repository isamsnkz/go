package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/isamsnkz/go/config"
	"github.com/isamsnkz/go/model"
)

func CreateClan(w http.ResponseWriter, r *http.Request) {
	var clans []model.Clan
	if err := json.NewDecoder(r.Body).Decode(&clans); err != nil {
		http.Error(w, "Format data tidak valid", http.StatusBadRequest)
		return
	}

	if len(clans) == 0 {
		http.Error(w, "Daftar clan tidak boleh kosong", http.StatusBadRequest)
		return
	}

	var savedClans []model.Clan
	for _, clan := range clans {
		if err := config.DB.Create(&clan).Error; err != nil {
			for _, c := range savedClans {
				config.DB.Delete(&c)
			}
			http.Error(w, "Gagal menyimpan data. Mungkin ada tag yang duplikat.", http.StatusBadRequest)
			return
		}
		savedClans = append(savedClans, clan)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedClans)
}

func GetClans(w http.ResponseWriter, r *http.Request) {
	var clans []model.Clan
	config.DB.Find(&clans)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clans)
}

func GetClanByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var clan model.Clan
	if err := config.DB.First(&clan, id).Error; err != nil {
		http.Error(w, "Clan tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clan)
}

func UpdateClan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var clan model.Clan
	if err := config.DB.First(&clan, id).Error; err != nil {
		http.Error(w, "Clan tidak ditemukan", http.StatusNotFound)
		return
	}

	var input model.Clan
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON tidak valid", http.StatusBadRequest)
		return
	}

	if input.ClanTags == "" || input.ClanName == "" || input.ClanType == "" || input.ClanLocation == "" {
		http.Error(w, "Semua field wajib diisi", http.StatusBadRequest)
		return
	}

	clan.ClanTags = input.ClanTags
	clan.ClanName = input.ClanName
	clan.ClanType = input.ClanType
	clan.ClanLocation = input.ClanLocation

	if err := config.DB.Save(&clan).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			http.Error(w, "ClanTags sudah digunakan", http.StatusBadRequest)
			return
		}
		http.Error(w, "Gagal memperbarui data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clan)
}

func DeleteClan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var clan model.Clan
	if err := config.DB.First(&clan, id).Error; err != nil {
		http.Error(w, "Clan tidak ditemukan", http.StatusNotFound)
		return
	}
	config.DB.Delete(&clan)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Clan berhasil dihapus"})
}
