// Souběžnost a paralelnost
//
// Seznam příkladů uložených v tomto podadresáři:
//
// 01_01_gorutine.go:
// - vytvoření a zavolání dvou gorutin
//
// 01_02_gorutine_wait.go:
// - vytvoření a zavolání dvou gorutin
// - naivní čekání na dokončení gorutin
//
// 01_03_goroutines_3.go:
// - vytvoření a zavolání různých gorutin
// - gorutiny budou zapisovat znaky na stejný textový řádek
//
// 01_04_goroutines_4.go:
// - vytvoření a zavolání různých gorutin
// - gorutiny se volají z jiných gorutin
// - gorutiny budou zapisovat znaky na stejný textový řádek
//
// 02_01_channels.go:
// - použití kanálu pro čekání na dokončení gorutiny
//
// 02_02_channels.go:
// - vytvoření a zavolání jedné gorutiny
// - použití kanálu pro čekání na dokončení gorutiny
//
// 02_03_channels.go:
// - vytvoření a zavolání tří gorutin
// - použití kanálu pro čekání na dokončení gorutiny
//
// 02_04_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění workera
// - předání práce workerovi
// - nečeká se na dokončení workera
//
// 02_05_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění workera
// - předání práce workerovi
// - uzavření kanálu indikuje, že se má worker ukončit
//
// 02_06_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit
//
// 02_07_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit
// - aktivní čekání na dokončení workerů
//
// 02_08_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit
// - aktivní čekání na dokončení workerů
//
// 02_09_worker.go:
// - funkce s implementací workera
// - asynchronní spuštění tří workerů
// - předání práce workerovi
// - uzavření kanálu indikuje, že se mají workeři ukončit
// - aktivní čekání na dokončení workerů
// - konstrukce for-range s kanálem
//
// 02_10_deadlock.go:
// - deadlock při práci s kanály
// - čekání na prvek z kanálu, do kterého žádná gorutina neprovádí zápis
//
// 02_11_deadlock.go:
// - deadlock při práci s kanály
// - zápis do kanálu, ze kterého žádná gorutina neprovádí čtení
//
// 03_01_three_channels.go:
// - konstrukce select-case
// - čekání na data poslaná do libovolného kanálu
//
// 03_02_three_channels.go:
// - konstrukce select-case
// - čekání na data poslaná do libovolného kanálu
//
// 03_03_select_default.go:
// - konstrukce select-case
// - věvev default v konstrukci select-case
//
// 03_04_select.go:
// - konstrukce select-case
// - explicitní ukončení workera

