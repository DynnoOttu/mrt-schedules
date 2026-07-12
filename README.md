# MRT Schedules API

REST API sederhana berbasis Go (Gin framework) untuk mengecek daftar stasiun dan jadwal keberangkatan MRT Jakarta.

Project ini dibuat sebagai bahan belajar Go dan Gin, mengambil data dari API publik MRT Jakarta (`https://www.jakartamrt.co.id/id/val/stasiuns`). Karena API tersebut saat ini sudah tidak mengembalikan data JSON (kemungkinan sudah tidak aktif/berubah), aplikasi ini menyediakan **fallback data dummy** agar tetap bisa dijalankan dan dipelajari.

## Tech Stack

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- `encoding/json` — parsing response API

## Fitur

- Mendapatkan daftar seluruh stasiun MRT
- Mengecek jadwal keberangkatan berdasarkan ID stasiun (arah Bundaran HI & Lebak Bulus)
- Fallback otomatis ke data dummy jika API sumber sedang tidak bisa diakses

## Instalasi & Menjalankan

1. Clone repository
   ```bash
   git clone https://github.com/DynnoOttu/mrt-schedules.git
   cd mrt-schedules
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Jalankan aplikasi
   ```bash
   go run main.go
   ```

Server akan berjalan di `http://localhost:8080` (sesuaikan port jika berbeda).

## API Endpoints

### 1. Get All Stations

Mengambil daftar seluruh stasiun MRT.

**Request**
```
GET /api/station
```

**Response**
```json
{
    "success": true,
    "message": "Successfully get all station",
    "data": [
        {
            "id": "1",
            "name": "Bundaran HI"
        },
        {
            "id": "2",
            "name": "Dukuh Atas BNI"
        },
        {
            "id": "3",
            "name": "Setiabudi Astra"
        }
    ]
}
```

### 2. Get Schedule by Station

Mengambil jadwal keberangkatan (arah Bundaran HI dan Lebak Bulus) berdasarkan ID stasiun.

**Request**
```
GET /api/station/:id
```

| Parameter | Tipe   | Keterangan          |
|-----------|--------|---------------------|
| `id`      | string | ID/NID dari stasiun |

**Response — Berhasil**
```json
{
    "success": true,
    "message": "Successfully get schedule by station",
    "data": [
        {
            "station": "Bundaran HI",
            "time": "05:00"
        },
        {
            "station": "Bundaran HI",
            "time": "05:30"
        },
        {
            "station": "Bundaran HI",
            "time": "06:00"
        },
        {
            "station": "Bundaran HI",
            "time": "05:10"
        },
        {
            "station": "Bundaran HI",
            "time": "05:40"
        },
        {
            "station": "Bundaran HI",
            "time": "06:10"
        }
    ]
}
```

**Response — Stasiun Tidak Ditemukan**
```json
{
    "success": false,
    "message": "station not found",
    "data": null
}
```

## Catatan

API sumber data asli (`jakartamrt.co.id`) saat ini sudah tidak mengembalikan JSON yang valid, sehingga aplikasi ini menggunakan data dummy sebagai fallback ketika request atau parsing ke API asli gagal. Struktur data dummy dibuat mengikuti format asli API (`nid`, `title`, `jadwal_hi_biasa`, `jadwal_lb_biasa`) agar mudah dikembalikan ke sumber asli jika API tersebut aktif kembali di kemudian hari.

## Struktur Project

```
mrt-schedules/
├── commont/
│   ├── client/
│   │   └── client.go        # helper HTTP client (DoRequest)
│   └── response/
│       └── response.go      # struct APIResponse (format response seragam)
├── modules/
│   └── stations/
│       ├── dto.go           # struct Station, StationResponse, Schedule, ScheduleResponse
│       ├── router.go        # routing endpoint (GET /api/station, /api/station/:id)
│       └── service.go       # business logic (GetAllStations, CheckScheduleByStation)
├── go.mod
└── main.go
```

## Author

Dynno — [github.com/DynnoOttu](https://github.com/DynnoOttu)
