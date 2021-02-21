#2 Tipe Package
1. Executable -> file yang akan dieksekusi. Harus menggunakan main
2. Liblary -> Dipanggil di file lain untuk diekskusi

#Jenis Import
1. Import standart lib golang
2. Package yang kita bikin sendiri
3. Dibikin orang lain

#Variabel Definition

var "variabel_name" "data_type" -> Default Value
var "variabel_name" "data_type" = "isinya" -> With assigment

"variabel_name" := "isinya"

#Default Value
String -> ""
Integer -> 0

#MAP
var "var_name" map["key_type"]"data_type"
delet("map_object","map_key") -> delete data in map
"value_variable","is_eixst_variabel" := "map_name"["map_key"]

#Funcition
func calculate(p, l int) (int, int) 
func calculate(p, l int) (luas int, keliling int) 