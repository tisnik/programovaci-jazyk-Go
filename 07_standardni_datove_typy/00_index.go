// Základní datové typy v jazyce Go
//
// Seznam příkladů uložených v tomto podadresáři:
//
// 01_01_standard_types_identifiers.go:
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
//
// 01_02_zero_values.go:
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk výchozích (nulových) hodnot těchto proměnných
//
// 01_03_standard_types_print.go:
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk jmen typů proměnných
//
// 01_04_types_and_zero_values.go:
// - deklarace proměnných, z nichž každá je odlišného
//   základního datového typu (nebo jeho aliasu)
// - tisk jmen typů proměnných i jejich (nulových) hodnot
//
// 02_01_integer_global_constant.go:
// - celočíselná hodnota deklarovaná jako globální konstanta
//
// 02_02_integer_local_constant.go:
// - celočíselná hodnota deklarovaná jako lokální konstanta
//
// 02_03_integer_constant_type.go:
// - celočíselná hodnota deklarovaná jako lokální konstanta
// - tisk typu této konstanty
//
// 02_04_integer_signed_types.go:
// - využití celočíselných datových typů se znaménkem
//
// 02_05_integer_signed_types_checks.go:
// - kontroly prováděné překladačem při přiřazování celočíselných konstant
//   do proměnných
//
// 02_06_integer_unsigned_types.go:
// - využití celočíselných datových typů bez znaménka
//
// 02_07_integer_constants.go:
// - celočíselné konstanty zapisované ve dvojkové, osmičkové a šestnáctkové
//   (hexadecimální) soustavě
//
// 02_08_integer_from_constant.go:
// - lokální proměnné typu celé číslo se znaménkem
// - inicializace lokálních proměnných globální konstantou
//
// 02_09_integer_from_constant.go:
// - lokální proměnné typu celé číslo bez znaménka
// - inicializace lokálních proměnných globální konstantou
//
// 02_10_integer_underscores.go:
// - celočíselné konstanty zapisované ve dvojkové, osmičkové a šestnáctkové
//   (hexadecimální) soustavě
// - použití znaku _ pro vizuální oddělení číslic
//
// 02_11_int_uint.go:
// - datové typy int a uint
//
// 02_12_improper_conversion.go:
// - striktní kontroly prováděné překladačem
//   při konverzi mezi datovými typy
//
// 02_13_improper_conversion.go:
// - striktní kontroly prováděné překladačem
//   při konverzi mezi datovými typy
//
// 02_14_improper_conversion_int_uint.go:
// - striktní kontroly prováděné překladačem při konverzi
//   mezi datovými typy
//
// 02_15_explicit_conversions.go:
// - explicitní konverze mezi hodnotami různých datových typů
//
// 02_16_formatting_output.go:
// - tisk celočíselných numerických hodnot s jejich naformátováním
//
// 02_17_formatting_output.go:
// - tisk celočíselných numerických hodnot s jejich naformátováním
//
// 02_18_int_subtraction.go:
// - rozdíl dvou hodnot typu int
//
// 02_19_int_subtration.go:
// - rozdíl dvou hodnot typu unsigned int
//
// 03_01_fp_types.go:
// - datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
//   float32 a float64.
//
// 03_02_fp_types_checks.go:
// - kontroly prováděné překladačem, zda numerické hodnoty s plovoucí čárkou
//   nepřekračují povolený rozsah.
//
// 03_03_fp_constants.go:
// - datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
//   float32 a float64
// - zápis literálů s numerickými hodnotami.
//
// 03_04_fp_hex_constants.go:
// - datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
//   float32 a float64
// - hexadecimální mantisa a exponent.
//
// 03_05_fp_special_values.go:
// - datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
//   float32 a float64
// - speciální hodnoty těchto typů
//
// 03_06_fp_global_constant.go:
// - hodnota s plovoucí čárkou deklarovaná jako globální konstanta
//
// 03_07_fp_local_constant.go:
// - hodnota s plovoucí čárkou deklarovaná jako lokální konstanta
//
// 03_08_fp_constant_type.go:
// - hodnota s plovoucí čárkou deklarovaná jako lokální konstanta
// - tisk typu této konstanty
//
// 03_09_fp_constant_type.go:
// - hodnota s plovoucí čárkou deklarovaná jako lokální konstanta
// - explicitní převod na typ float32
// - tisk typu této konstanty
//
// 03_10_fp_print_values.go:
// - tisk numerických hodnot s plovoucí řádovou čárkou
//   s jejich naformátováním.
//
// 03_11_fp_print_values.go:
// - tisk numerických hodnot s plovoucí řádovou čárkou
//   s jejich naformátováním.
//
// 04_01_complex_types.go:
// - deklarace proměnných typu komplexní číslo
//
// 04_02_complex_global_constant.go:
// - komplexní hodnota deklarovaná jako globální konstanta
// - tisk této hodnoty
//
// 04_03_complex_local_constant.go:
// - komplexní hodnota deklarovaná jako lokální konstanta
// - tisk této hodnoty
//
// 04_04_complex_constant_type.go:
// - komplexní hodnota deklarovaná jako lokální konstanta
// - tisk typu této konstanty
//
// 04_05_complex64_real_imag.go:
// - získání reálné a imaginární složky komplexního čísla
// - varianta programu pro typ complex64
//
// 04_06_complex128_real_imag.go:
// - získání reálné a imaginární složky komplexního čísla
// - varianta programu pro typ complex128
//
// 04_07_complex_arithmetic.go:
// - výpočty s komplexními čísly
//
// 04_08_complex_modulo.go:
// - pokus o provedení operace modulo
//
// 04_09_complex_special_values.go:
// - speciální komplexní hodnoty z balíčku math/cmplx
//
// 04_10_complex_special_values.go:
// - speciální komplexní hodnoty získané jako výsledky výpočtů
//
// 05_01_boolean_type.go:
// - datový typ bool a pravdivostní hodnoty
//
// 05_02_boolean_type_checks.go:
// - datový typ bool a pravdivostní hodnoty v programovacím jazyku Go
// - překladač jazyka Go neumožní automatickou konverzi mezi nějakým
//   datovým typem odlišným od bool na pravdivostní hodnotu.
//
// 05_03_boolean_global_constant.go:
// - pravdivostní hodnota deklarovaná jako globální konstanta
//
// 05_04_boolean_local_constant.go:
// - pravdivostní hodnota deklarovaná jako lokální konstanta
//
// 05_05_boolean_constant_type.go:
// - pravdivostní hodnota deklarovaná jako lokální konstanta
// - tisk typu konstanty (ne její hodnoty)
//
// 05_06_boolean_print.go:
// - tisk hodnoty typu boolean
// - využití různých typů formátu
//
