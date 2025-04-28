// Složené datové typy
//
// Seznam příkladů uložených v tomto podadresáři:
//
// 01_01_string_global_constant.go:
// - deklarace globální konstanty s řetězcem
//
// 01_02_string_local_constant.go:
// - deklarace lokální konstanty s řetězcem
//
// 01_03_string_constant_type.go:
// - deklarace konstanty s řetězcem
// - tisk typu této konstanty
//
// 01_04_string_variable.go:
// - deklarace proměnné typu řetězec
// - tisk obsahu této proměnné
//
// 01_05_string_type_inference.go:
// - deklarace proměnné typu řetězec
// - využití typové inference
//
// 01_06_string_type.go:
// - deklarace proměnné typu řetězec
// - využití typové inference
// - tisk typu proměnné
//
// 01_07_strings.go:
// - prázdný řetězec
// - řetězce s řídicími znaky
// - řetězce s Unicode znaky
//
// 01_08_raw_strings.go:
// - deklarace a použití takzvaných surových řetězců
//   (raw strings)
//
// 01_09_multiline_string.go:
// - deklarace víceřádkového řetězce
//
// 01_10_multiline_string.go:
// - deklarace víceřádkového řetězce
//
// 01_11_string_length.go:
// - zjištění délky řetězce v bajtech
//
// 01_12_string_for_loop_ascii.go:
// - řetězec obsahující jen ASCII znaky
// - průchod jednotlivými bajty tvořícími řetězec
//
// 01_13_string_for_loop_unicode.go:
// - řetězec obsahující Unicode znaky
// - průchod jednotlivými bajty tvořícími řetězec
//
// 01_14_string_for_loop_unicode.go:
// - výpis hexadecimálních hodnot bajtů tvořících řetězec
//
// 01_15_string_range.go:
// - řetězec s Unicode znaky
// - průchod jednotlivými znaky řetězce
// - tisk kódů znaků
//
// 01_16_string_range.go:
// - řetězec s Unicode znaky
// - průchod jednotlivými znaky řetězce
// - tisk těchto znaků
//
// 01_17_string_length_chars.go:
// - zjištění délky řetězce v runách
//
// 01_18_strings.go:
// - prázdný řetězec
// - řetězce s řídicími znaky
// - řetězce s Unicode znaky
//
// 02_01_array_global_variable.go:
// - deklarace globální proměnné s polem
//
// 02_02_array_local_variable.go:
// - deklarace lokální proměnné s polem
//
// 02_03_array_initialization.go:
// - deklarace lokální proměnné s polem
// - inicializace pole
//
// 02_04_array_initialization.go:
// - deklarace lokální proměnné s polem
// - inicializace pole
// - počet prvků si zjistí překladač automaticky
//
// 02_05_array_partial_initialization.go:
// - deklarace lokální proměnné s polem
// - inicializace pole
// - prvky bez explicitně uvedené hodnoty jsou vynulovány
//
// 02_06_array_type_inference.go:
// - deklarace lokální proměnné s polem
// - použití typové inference
// - inicializace pole
//
// 02_07_array_type_inference.go:
// - deklarace lokální proměnné s polem
// - použití typové inference
// - inicializace pole
//
// 02_08_array_type_inference.go:
// - deklarace lokální proměnné s polem
// - použití typové inference
// - inicializace pole
// - odvození délky pole překladačem
//
// 02_09_array_length.go:
// - zjištění délky pole (počtu prvků)
//
// 02_10_array_capacity.go:
// - zjištění kapacity pole (počtu prvků)
//
// 02_11_array_modification.go:
// - modifikace pole prvek po prvku
//
// 02_12_array_modification.go:
// - modifikace pole prvek po prvku
//
// 02_13_array_copy.go:
// - kopie pole
// - modifikace pole, ale nikoli jeho kopie
//
// 02_14_matrix.go:
// - dvourozměrná pole - matice
//
// 02_15_matrix.go:
// - dvourozměrná pole - matice
//
// 02_16_pass_array.go:
// - předání pole do funkce
// - modifikace pole v této funkci
// - zjištění, zda se pole předává hodnotou nebo odkazem
//
// 02_17_pass_array.go:
// - předání pole do funkce přes ukazatel (odkaz)
// - modifikace pole v této funkci
// - zjištění, zda se pole předává hodnotou nebo odkazem
//
// 02_18_array_for_loop.go:
// - průchod všemi prvky pole
//
// 02_19_array_for_loop.go:
// - průchod všemi prvky pole
//
// 03_01_new_nil_slice.go:
// - konstrukce takzvaného nulového řezu (nil slice)
//
// 03_02_new_empty_slice.go:
// - konstrukce prázdného řezu
//
// 03_03_make_slice.go:
// - konstrukce prázdného řezu pomocí make
// - řez bude mít nulovou kapacitu
//
// 03_04_make_slice.go:
// - konstrukce prázdného řezu pomocí make
// - řez bude mít nenulovou kapacitu
//
// 03_05_modify_nil_slice.go:
// - konstrukce nulového řezu
// - pokus o modifikaci prvního prvku, který neexistuje
//
// 03_06_modify_filled_slice.go:
// - konstrukce řezu s jeho inicializací
// - modifikace prvního prvku uloženého v řezu
//
// 03_07_modify_filled_slice.go:
// - konstrukce řezu s jeho inicializací
// - modifikace prvního prvku uloženého v řezu
//
// 03_08_slice_vs_array.go:
// - rozdíl mezi řezem a polem
//
// 03_09_slice_vs_array.go:
// - rozdíl mezi řezem a polem
//
// 03_10_slice_from_array.go:
// - vytvoření řezů z pole
//
// 03_11_array_slice_modification.go:
// - modifikace řezu
// - zjištění, jak se změní pole, na které řez ukazuje
//
// 03_12_slice_from_slice.go:
// - vytvoření řezu z jiného řezu
//
// 03_13_slice_append.go:
// - postupné přidávání prvků do řezu
// - ve skutečnosti se vždy vytvoří nový řez
//
// 03_14_make_slice.go:
// - konstrukce prázdného řezu pomocí make
// - řez bude mít nenulovou kapacitu
//
// 03_15_slice_modification.go:
// - vytvoření řezu z jiného řezu
// - zjištění, jak se modifikace prvku v jednom řezu
//   projeví na druhém řezu
//
// 03_16_slice_range_1.go:
// - vytvoření řezu
// - průchod všemi prvky řezu
//
// 03_17_slice_range_2.go:
// - vytvoření řezu
// - průchod všemi prvky řezu
//
// 04_01_uninitialized_map.go:
// - deklarace nulové mapy
// - pokus o přidání dvojice klíč-hodnota do mapy
//
// 04_02_initialized_map.go:
// - deklarace prázdné mapy
// - přidání dvojic klíč-hodnota do mapy
//
// 04_03_initialized_map.go:
// - zkrácená deklarace prázdné mapy
// - přidání dvojic klíč-hodnota do mapy
//
// 04_04_preallocated_map.go:
// - zkrácená deklarace prázdné mapy
// - mapa má předalokovanou kapacitu
// - přidání dvojic klíč-hodnota do mapy
//
// 04_05_dictionary.go:
// - klasický slovník: klíče i hodnoty jsou řetězci
//
// 04_06_map_capacity.go:
// - získání délky mapy
//
// 04_07_reading_from_maps.go:
// - přečtení prvku se známým klíčem
// - přečtení prvku s neznámým klíčem
//
// 04_08_reading_from_maps.go:
// - přečtení prvku se známým klíčem s testem existence
// - přečtení prvku s neznámým klíčem s testem existence
//
// 04_09_delete_from_map.go:
// - smazání existujícího prvku z mapy
// - smazání neexistujícího prvku z mapy
//
// 04_10_map_range_1.go:
// - průchod všemi prvky mapy
//
// 04_11_map_range_2.go:
// - průchod všemi hodnotami prvků mapy
//
// 04_12_map_range_3.go:
// - průchod všemi klíči mapy
//
// 05_01_struct.go:
// - deklarace uživatelské datové struktury
// - zápis hodnot prvků do datové struktury
//
// 05_02_struct_init.go:
// - inicializace prvků datové struktury
// - zápis inicializace odvozený od jazyka C
//
// 05_03_better_struct_init.go:
// - inicializace prvků datové struktury
// - vylepšený zápis inicializace s vyjmenováním prvků
//
// 05_04_struct_comparison.go:
// - inicializace prvků datové struktury
// - vylepšený zápis inicializace s vyjmenováním prvků
// - porovnání dvou struktur
//
// 05_05_array_of_structs.go:
// - pole, jehož prvky jsou typu struktura (záznam)
//
// 05_06_array_of_structs.go:
// - pole, jehož prvky jsou typu struktura (záznam)
// - přímá inicializace prvků pole
//
// 05_07_map_and_struct.go:
// - mapa, která obsahuje uložené záznamy
// - jako klíče jsou použity řetězce
//
// 05_08_map_and_struct.go:
// - mapa, která obsahuje uložené záznamy
// - jako klíče jsou použity taktéž záznamy
//
// 05_09_struct_func.go:
// - funkce, která jako parametr akceptuje záznam
//
// 05_10_struct_func.go:
// - funkce, která jako parametr akceptuje záznam
// - funkce, která vytvoří záznam
//
// 06_01_nil.go:
// - proměnné různých typů, které jsou automaticky
//   nastaveny na hodnotu nil
//
