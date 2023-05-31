package seeder

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"tamiyochi-backend/entity"
// 	"time"

// 	"github.com/gocarina/gocsv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // Entry defines both the CSV layout and database schema
// type Entry struct {
//     // gorm.Model

//     Field1 string `csv:"title"`
//     Field2 string `csv:"synopsis"`
//     Field3 string `csv:"start_date"`
//     Field4 string `csv:"score"`
//     Field5 string `csv:"scored_by"`
//     Field6 string `csv:"members"`
//     Field7 string `csv:"main_picture"`
//     Field8 string `csv:"serializations"`
//     Field9 int    `csv:"manga_id"`
//     Field10 string `csv:"genres"`
//     Field11 string `csv:"authors"`
// }

// type Iot struct {
//     Id              int             `json:"id"`
//     FirstName       string          `json:"first_name"`
//     LastName        string          `json:"last_name"` // RawMessage here! (not a string)
//     Role            string          `json:"role"`
// }

// func main() {
//     // Open the CSV file for reading
//     file, err := os.Open("manga_mal.csv")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()

//     var entries []Entry
//     err = gocsv.Unmarshal(file, &entries)
//     if err != nil {
//         panic(err)
//     }

//     dsn := fmt.Sprintf("host=localhost user=postgres password=rencist dbname=tamiyochi port=5432 TimeZone=Asia/Jakarta")
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(db)
// 		panic(err)
// 	}
//     arrHarga := [8]int{18000, 20000, 15000, 24000, 22000, 16000, 25000, 30000}
//     idCount := 1;
//     for _, entry := range entries {
//         var manga entity.Manga

//         rand.Seed(time.Now().UnixNano())
//         mangaCount := 1 + rand.Intn(8-1+1)
//         for i := 1; i < mangaCount+1; i++ {
//             manga.ID = idCount
//             idCount++

//             manga.SeriID = entry.Field9
//             manga.Volume = i

//             rand.Seed(time.Now().UnixNano())
//             volumeCount := 1 + rand.Intn(5-1+1)
//             manga.JumlahTersedia = volumeCount

//             rand.Seed(time.Now().UnixNano())
//             hargaCount := 0 + rand.Intn(7-0+1)
//             manga.HargaSewa = arrHarga[hargaCount]

//             db.Create(&manga)
//         }
//     }

//     // for _, entry := range entries {
//     //     // mangaID := entries[0].Field9
//     //     genre := entry.Field10
//     //     genreArray := strings.Split(genre, ",")
//     //     for _, lmao := range genreArray {
//     //         var genre entity.Genre
//     //         if string(lmao[1]) == "]" {
//     //             rand.Seed(time.Now().UnixNano())
//     //             n := 1 + rand.Intn(225-1+1)
//     //             genre.ID = n
//     //         } else {
//     //             lmao2 := lmao[2:]
//     //             lmao3 := lmao2[:len(lmao2)-1]
//     //             lmaoFinal := strings.Replace(lmao3, "'", "", 1)

//     //             db.Where("nama = ?", lmaoFinal).First(&genre)
//     //         }
//     //         var seriGenre = entity.SeriGenre{
//     //             SeriID: entry.Field9,
//     //             GenreID: genre.ID,
//     //         }
//     //         db.Create(&seriGenre)
//     //     }

//     //     // rand.Seed(time.Now().UnixNano())
//     //     // n := 1 + rand.Intn(225-1+1)
//     //     // if seri.PenerbitID == 31 {
//     //     //     seri.PenerbitID = n
//     //     // }

//     //     // var seriPenerbit = entity.P{
//     //     //     SeriID: seri.ID,
//     //     //     PenerbitID: seri.PenerbitID,
//     //     // }
//     //     // db.Create(&seri)
//     // }

//     // for _, entry := range entries {
//     //     var penerbit = entity.Penerbit{}
//     //     db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2)).First(&penerbit)
//     //     var seri = entity.Seri{
//     //         ID: entry.Field9,
//     //         Judul: entry.Field1,
//     //         Sinopsis: entry.Field2,
//     //         TahunTerbit: entry.Field3,
//     //         Skor: entry.Field4,
//     //         TotalPenilai: entry.Field5,
//     //         TotalPembaca: entry.Field6,
//     //         Foto: entry.Field7,
//     //         PenerbitID: penerbit.ID,
//     //     }

//     //     rand.Seed(time.Now().UnixNano())
//     //     n := 1 + rand.Intn(225-1+1)
//     //     if seri.PenerbitID == 31 {
//     //         seri.PenerbitID = n
//     //     }
//     //     db.Create(&seri)
//     // }

//     // fmt.Println(entries[0].Field1 + " | " + entries[0].Field2 + " | " + entries[0].Field3 + " | " + entries[0].Field4 + " | " + entries[0].Field5 + " | " + entries[0].Field6 + " | " + entries[0].Field7 + " | " + entries[0].Field8)
//     // for _, entry := range entries {

//     //     // if strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2) == "" {
//     //     //     fmt.Println(strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2))
//     //     // }

//     // }

//     // var iot Iot
//     // // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 10), "},")[1], " ", "`", 1))
//     // // for i := 0; i < 2; i++ {
//     // err = json.Unmarshal([]byte(`{"id": 1868, "first_name": "Kentarou", "last_name": "Miura", "role": "Story & Art"}`), &iot)
//     // if err != nil {
//     //     panic(err)
//     // }
//     // // }
//     // // Context is []byte, so you can keep it as string in DB
//     // fmt.Println(iot.Id)

//     // genre := [17]string{"Action",
//     // "Adventure",
//     // "AwardWinning",
//     // "Drama",
//     // "Fantasy",
//     // "Horror",
//     // "Supernatural",
//     // "Mystery",
//     // "SliceOfLife",
//     // "Psychological",
//     // "Comedy",
//     // "Suspense",
//     // "Sports",
//     // "Sci-Fi",
//     // "Ecchi",
//     // "Romance",
//     // "GirlsLove"}

//     // for q := 0; q < 2; q++ {
//     //     for i := 0; i < strings.Count(entries[q].Field1, ",") + 1; i++ {
//     //         fmt.Println(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Split(entries[q].Field1, ",")[i],"[","", 1), """, "", 10), "]", "", 1), " ", "", 1) )
//     //     }
//     //     fmt.Println(" |||| ")
//     //     // fmt.Println(strings.Replace(strings.Replace(strings.Replace(strings.Split(entry.Field1, ",")[0],"[","", 1), "'", "", 2), "]", "", 1))
//     //     // fmt.Println(entry.Field1)
//     // }

// 	// dsn := fmt.Sprintf("host=localhost user=postgres password=rencist dbname=tamiyochi port=5432 TimeZone=Asia/Jakarta")
// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	panic(err)
// 	// }
//     // count := 1
//     // for _, entry := range entries {
//         // var penerbitEntity entity.Penerbit
//         // dbRresult := db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2)).First(&penerbitEntity)
//     //     if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
//     //         penerbit := entity.Penerbit{
//     //             ID: count,
//     //             Nama: strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2),
//     //         }
//     //         count++
//     //         db.Create(&penerbit)
//     //     }
//         // var penerbitEntity entity.Penerbit
// 	    // db.Find(&penerbitEntity)
//         // db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2)).Take(&penerbitEntity)
//         // // if err != nil {
//         //     if penerbitEntity.Nama != strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2) {
//         //         penerbit := entity.Penerbit{
//         //             ID: count,
//         //             Nama: strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2),
//         //         }
//         //         count++
//         //         db.Create(&penerbit)
//         //     }
//         // } else {
//             // fmt.Println(err)
//         // }

//         // -------------------------------------------------------------------------------------------------
//         // var penulis entity.Penulis
//         // count1 := 1
//         // for _, entry := range entries {
//         //     count := strings.Count(entry.Field11, "{")
//         //     if count > 1 {
//         //         var penulisSeri entity.PenulisSeri
//         //         var iot Iot
//         //         err = json.Unmarshal([]byte(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100), "},")[1], " ", "", 1)), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.LastName)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // db.Create(&penulis)

//         //         err = json.Unmarshal([]byte(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100), ", {")[0]), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.LastName)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //         // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         //         // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         //     } else {
//         //         var penulisSeri entity.PenulisSeri
//         //         var iot Iot
//         //         err = json.Unmarshal([]byte(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100)), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.Id)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // db.Create(&penulis)
//         //         // fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //         // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100))
//         //     }
//         // }
//         // -------------------------------------------------------------------------------------------------

//         // count := strings.Count(entries[i].Field1, "{")
//         // if count > 1 {
//         //     var iot Iot
//         //     err = json.Unmarshal([]byte(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100), "},")[1], " ", "", 1)), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.LastName)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     err = json.Unmarshal([]byte(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100), ", {")[0]), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.LastName)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         //     // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         // } else {
//         //     var iot Iot
//         //     err = json.Unmarshal([]byte(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100)), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.Id)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100))
//         // }

//         // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entries[1].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100))

//     }
