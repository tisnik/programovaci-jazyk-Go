// Základy programovacího jazyka Go
//
// Seznam příkladů uložených v tomto podadresáři:
//
// 01_01_keywords_as_identifiers.go:
// - pokus o použití klíčových slov ve funkci identifikátorů
//
// 01_02_int_type_redeclaration.go:
// - deklarace datového typu nazvaného "int", který přepíše původní typ "int"
//        a především, kdy ho naopak nepoužít.
//
// 01_03_true_redeclaration.go:
// - deklarace proměnné "true", která přepíše původní hodnotu "true"
//        a především, kdy ho naopak nepoužít.
//
// 02_01_unary_operators.go:
// - unární operátory definované v jazyku Go
//
// 02_02_arithmetic_operators_int.go:
// - základní aritmetické operátory použité s proměnnými
//   typu int
//
// 02_03_arithmetic_operators_float.go:
// - základní aritmetické operátory s proměnnými
//   typu float32
//
// 02_04_arithmetic_operators_complex.go:
// - základní aritmetické operátory s proměnnými
//   typu complex64
//
// 02_05_div_mod.go:
// - operátory pro celočíselný podíl a výpočet zbytku po dělení
// - použití pro různé kombinace kladných i záporných operandů
//
// 02_06_relational_operators_int.go:
// - základní relační operátory s proměnnými
//   typu int
//
// 02_07_relational_operators_float.go:
// - základní relační operátory s proměnnými
//   typu float32
//
// 02_08_relational_operators_complex.go:
// - základní relační operátory s proměnnými
//   typu complex64
//
// 02_09_relational_operators_boolean.go:
// - základní relační operátory s proměnnými
//   typu boolean
//
// 02_10_relational_operators_strings.go:
// - základní relační operátory s proměnnými
//   typu řetězec
//
// 02_11_logical_operators.go:
// - logické operátory (pro hodnoty typu bool)
//
// 02_12_bitwise_operators_8b.go:
// - bitové operátory pro osmibitové hodnoty
//
// 02_13_bitwise_operators_32b.go:
// - bitové operátory pro 32bitové hodnoty
//
// 02_14_bit_shifts_byte.go:
// - bitové posuny pro hodnoty typu byte
//
// 02_15_bit_shifts_int.go:
// - bitové posuny pro hodnoty typu int
//
// 02_16_bit_shifts.go:
// - bitové posuny o hodnotu proměnné (bez znaménka)
//
// 02_17_arithmetic_assignment.go:
// - základní aritmetické operátory kombinované
//   s přiřazením
//
// 02_18_bit_shift_assignment.go:
// - operátory pro bitové posuny kombinované
//   s přiřazením
//
// 02_19_bitops_assignment.go:
// - operátory pro bitové operace kombinované
//   s přiřazením (použito u proměnné typu byte)
// - použití celočíselných konstant zapsaných
//   ve dvojkové soustavě
//
// 02_20_inc_dec_int.go:
// - operátory ++ a -- použité pro proměnné typu int
//
// 02_21_inc_dec_float.go:
// - operátory ++ a -- použité pro proměnné typu float32
//
// 02_22_inc_dec_complex.go:
// - operátory ++ a -- použité pro proměnné typu complex
//
// 02_23_inc_dec_bad_usage_1.go:
// - nekorektní způsob použití operací ++ a --
//
// 02_24_inc_dec_bad_usage_2.go:
// - nekorektní způsob použití operací ++ a --
//
// 02_25_string_concatenation.go:
// - spojení řetězců operátorem +
//
// 02_26_short_circuit.go:
// - zkrácené vyhodnocení operátorů && a ||
//
// 04_01_const_1.go:
// - použití klíčového slova const
// - deklarace konstanty "a" globální pro celý modul
// - deklarace konstanty "b", která je platná jen v těle funkce "main"
//
// 04_02_const_2.go:
// - použití klíčového slova const
// - deklarace konstanty "a" globální pro celý modul
// - deklarace konstanty "b", která je platná jen v těle funkce "main"
// - deklarace konstanty "c", která je platná v jednom vnořeném bloku
// - deklarace konstanty "d", která je platná v dalším vnořeném bloku
//
// 04_03_const_3.go:
// - použití klíčového slova const
// - rozdíl mezi veřejnou (public) konstantou a konstantou soukromou (private)
//
// 04_04_const_4.go:
// - použití klíčového slova const
// - rozdíl mezi veřejnou (public) konstantou a konstantou soukromou (private)
// - zápis deklarace konstant do společného bloku
//
// 04_05_const_types.go:
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
//
// 04_06_const_types.go:
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - přiřazení pojmenovaných konstant do proměnných
// - použití pojmenovaných konstant ve výrazu
//
// 04_07_const_types.go:
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - pokus o přiřazení do proměnných typu MyString
//
// 04_08_const_types.go:
// - použití klíčového slova const
// - konstanta bez explicitně definovaného typu
// - konstanta s definovaným typem hodnoty
// - výpis typů hodnot konstant
// - výpis typů proměnných, do kterých je přiřazena konstanta
//
// 04_09_numeric_const.go:
// - použití klíčového slova const
// - základní vlastnosti numerických konstant
//
// 04_10_const_expression.go:
// - použití klíčového slova const
// - konstanty inicializované konstantím výrazem
//
// 05_01_global_var_1.go:
// - použití klíčového slova var
// - deklarace globální proměnné
// - modifikace globální proměnné ve funkci
//
// 05_02_global_var_2.go:
// - použití klíčového slova var
// - deklarace globální proměnné
// - proměnná má výchozí (nulovou) hodnotu
// - modifikace globální proměnné ve funkci
//
// 05_03_global_var_2.go:
// - použití klíčového slova var
// - deklarace globální proměnné
// - proměnná má výchozí (nulovou) hodnotu
// - proměnná je viditelná mimo rámec balíčku
//
// 05_04_global_vars.go:
// - použití klíčového slova var
// - deklarace dvou globálních proměnných v bloku var
//
// 05_05_global_vars.go:
// - použití klíčového slova var
// - deklarace dvou globálních proměnných na jediném řádku
//
// 05_06_local_var_1.go:
// - použití klíčového slova var
// - deklarace lokální proměnné v rámci funkce
//
// 05_07_local_var_2.go:
// - použití klíčového slova var
// - deklarace lokální proměnné v rámci bloku
//
// 05_08_var_visibility.go:
// - použití klíčového slova var
// - oblast viditelnosti proměnných
//
// 05_09_var_visibility_2.go:
// - použití klíčového slova var
// - oblast viditelnosti proměnných
// - doplnění komentářů o hodnoty, které se vypíšou po spuštění
//
// 06_01_type_inference.go:
// - automatické odvození typu proměnné
//
// 06_02_type_inferences.go:
// - automatické odvození typu proměnné
//
// 06_03_type_inference_values.go:
// - automatické odvození typu proměnné
// - explicitní přetypování hodnot před přiřazením do proměnné
//
// 06_04_global_variables.go:
// - automatické odvození typu proměnné
// - pokus o deklaraci globálních proměnných
//
// 06_05_global_variables.go:
// - automatické odvození typu globálních proměnných
//
// 06_06_global_variables.go:
// - automatické odvození typu globálních proměnných
// - použití bloku var
//
// 07_01_elementary_function.go:
// - deklarace funkce bez parametrů a bez návratové hodnoty
//
// 07_02_print_sum.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
//
// 07_03_print_sum.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
//
// 07_04_print_sum_incorrect.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
// - funkce je volána s nekorektními typy argumentů
//
// 07_05_local_var_incorrect.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
// - nekorektní pokus o deklaraci proměnné pojmenované stejně jako parametr
//
// 07_06_calc_sum.go:
// - deklarace funkce s parametry a s návratovou hodnotou
// - parametry stejného typu jsou deklarovány společně
//
// 07_07_calc_sum_incorrect.go:
// - deklarace funkce s parametry a s návratovou hodnotou
// - parametry stejného typu jsou deklarovány společně
// - funkce nevrací žádnou hodnotu, i když by měla
//
// 07_08_multiple_return.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
//
// 07_09_named_return.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - pojmenování návratových hodnot
// - parametry stejného typu jsou deklarovány společně
//
// 07_10_named_return_incorrect.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - pojmenování návratových hodnot
// - parametry stejného typu jsou deklarovány společně
// - nekorektní deklarace proměnných, které se jmenují stejně, jako parametry
//
// 07_11_multiple_return_ignore.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
// - ignorování jedné návratové hodnoty
//
// 07_12_multiple_return_ignore.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
// - ignorování jedné návratové hodnoty
//
// 07_13_multiple_return_ignore.go:
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
// - ignorování obou návratových hodnot
//
// 07_14_argument_count.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
// - funkce je volána s nekorektními počty argumentů
//
// 07_15_varargs.go:
// - deklarace funkce s proměnným počtem parametrů
// - funkce je volána s různými počty argumentů
//
// 07_16_local_var_incorrect_2.go:
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
// - nekorektní pokus o deklaraci proměnné pojmenované stejně jako parametr
//
// 07_17_func_as_variable.go:
// - deklarace anonymní funkce
// - přiřazení této funkce do globální proměnné
// - zavolání této funkce s předáním argumentů
//
// 07_18_inner_function.go:
// - deklarace funkce uvnitř jiné funkce
// - zavolání této funkce s předáním argumentů
//
// 07_19_closure.go:
// - deklarace uzávěru s parametrem
// - opakované volání uzávěru
//
// 07_20_closure.go:
// - deklarace uzávěru
// - využití uzávěru (jeho zavolání)
//
// 07_21_closure.go:
// - deklarace uzávěru
// - konstruktor uzávěru
// - využití uzávěru z jiné funkce
//
// 07_22_closure.go:
// - deklarace uzávěru
// - několik nezávislých čítačů
//
// 07_23_closure.go:
// - deklarace uzávěru
// - konstruktor uzávěru s parametry
// - využití uzávěru z jiné funkce
//
// 07_24_closure.go:
// - deklarace uzávěru
// - několik nezávislých čítačů
//
// 07_25_closure.go:
// - deklarace více uzávěrů, které sdílí společnou navázanou proměnnou
//
// 07_26_closure.go:
// - deklarace více uzávěrů, které sdílí společnou navázanou proměnnou
// - při konstrukci uzávěrů lze předat argumenty
//   s počáteční hodnotou a krokem
//
// 08_01_if_statement.go:
// - základní varianta rozhodovací konstrukce "if"
//
// 08_02_if_statement.go:
// - rozhodovací konstrukce s lokální proměnnou
//
// 08_03_if_statement_bad_type.go:
// - řídicí konstrukce if a hodnoty, které nejsou typu boolean.
//
// 08_04_if_statement_nil.go:
// - řídicí konstrukce if a hodnoty, které nejsou typu boolean
//
// 08_05_if_else.go:
// - řídicí konstrukce if-else
//
// 08_06_bad_syntax.go:
// - příklady špatné syntaxe konstrukce if-else.
//
// 08_07_switch.go:
// - příklad rozvětvení pomocí konstrukce switch
//
// 08_08_switch.go:
// - příklad rozvětvení pomocí konstrukce switch
// - deklarace lokální proměnné přímo v konstrukci switch
//
// 08_09_constants_in_switch.go:
// - řídicí konstrukce switch s více konstantami
//
// 08_10_switch_comparisons.go:
// - konstrukce switch s podmínkami
//
// 08_11_switch_expressions.go:
// - řídicí konstrukce switch s výrazy ve větvích case.
//
// 08_12_switch_expressions.go:
// - řídicí konstrukce switch s výrazy ve větvích case.
//
// 08_13_no_fallthrough.go:
// - řídicí konstrukce switch - pracuje jinak než v C!
// - automatické ukončení jednotlivých větví tak,
//   jakoby byl zapsán příkaz break
//
// 08_14_fallthrough.go:
// - řídicí konstrukce switch - simulace chování C.
//
// 08_15_switch_string.go:
// - řídicí konstrukce switch s vyhodnocovaným výrazem typu string.
//
// 08_16_goto_from_switch.go:
// - výskok z příkazu switch pomocí goto.
//
// 08_17_switch_intro.go:
// - prázdná konstrukce switch
// - umístění větve default v konstrukci switch
//
// 08_18_switch_no_constants.go:
// - rozeskok na základě celočíselných konstant
//
// 08_19_switch_string.go:
// - řídicí konstrukce switch s vyhodnocovaným výrazem typu string
// - několik konstant zapsaných ve stejné větvi case
//
// 08_20_no_fallthrough.go:
// - řídicí konstrukce switch a implicitní příkazy break
//
// 08_21_fallthrough.go:
// - řídicí konstrukce switch a fallthrough
//
// 09_01_simplest_for_loop.go:
// - nejjednodušší forma programové smyčky for
//
// 09_02_for_loop_with_condition.go:
// - programová smyčka for s podmínkou na začátku
//
// 09_03_c_like_loop.go:
// - programová smyčka odvozená od jazyka C
// - jako počitadlo je použita běžná lokální proměnná
//
// 09_04_c_like_loop.go:
// - programová smyčka odvozená od jazyka C
// - počitadlem je proměnná viditelná jen v bloku smyčky
//
// 09_05_c_like_loop_local_var.go:
// - programová smyčka odvozená od jazyka C
// - počitadlem je proměnná viditelná jen v bloku smyčky
// - POZOR: tento příklad nelze přeložit
//
// 09_06_for_range.go:
// - programová smyčka založená na klíčovém slovu range
//
// 09_07_nested_loops.go:
// - vnořené programové smyčky
// - zobrazení tabulky malé násobilky
//
// 09_08_nested_loops.go:
// - vnořené programové smyčky
// - zobrazení tabulky malé násobilky
//
// 09_09_break_statement.go:
// - příkaz break v programových smyčkách
//
// 09_10_continue_statement.go:
// - příkaz continue v programových smyčkách
//
// 09_11_incomplete_break_from_inner_loop.go:
// - vnořené programové smyčky
// - příkaz break pro výskok z vnitřní smyčky
//
// 09_12_break_from_inner_loop.go:
// - vnořené programové smyčky
// - příkaz break pro výskok z vnořených smyček
//
// 09_13_continue_in_inner_loop.go:
// - vnořené programové smyčky
// - příkaz continue pro pokračování další iteraci vnitřní
//
// 09_14_continue_in_inner_loop.go:
// - vnořené programové smyčky
// - příkaz continue pro pokračování další iterace vnější smyčky
//
// 09_15_goto_from_loop.go:
// - jednoduchá programová smyčka
// - výskok ze smyčky pomocí goto
//
// 09_16_goto_from_inner_loop.go:
// - vnořené programové smyčky
// - výskok z vnořené smyčky pomocí goto
//
// 09_17_wrong_goto.go:
// - neplatné použití goto pro skok mezi funkcemi
//
// 09_18_wrong_goto.go:
// - neplatné použití goto pro skok mezi bloky
//
// 10_01_defer_basic_usage.go:
// - nejjednodušší způsob použití konstrukce defer
// - funkce onFinish bude zavolána při opouštění funkce main
//
// 10_02_defer_func.go:
// - anonymní funkce s odloženým voláním
// - zápis této funkce v kulatých závorkách
//
// 10_03_defer_func.go:
// - anonymní funkce s odloženým voláním
// - zjednodušený zápis této funkce
//
// 10_04_defer_with_parameters.go:
//
// 10_05_more_defers.go:
// - větší množství konstrukcí defer
// - ověření pořadí volání
//
// 10_06_defer_arguments_evaluation.go:
// - ověření, kdy se vyhodnocují parametry předávané funkci
//   s odloženým zavoláním
//
// 10_07_defer_arguments_evaluation.go:
// - ověření, kdy se vyhodnocují parametry předávané funkci
//   s odloženým zavoláním
// - předávání řezů
//
// 10_08_defer_on_all_returns.go:
// - ověření, zda se bude funkce s odloženým voláním
//   volat i v případě, že je použito větší množství
//   konstrukcí return
//
// 10_09_defer_practical_usage.go:
// - praktické použití konstrukce defer
// - realizace kopie obsahu souboru
// - delší varianta vypisující prováděné operace
//
// 10_10_defer_practical_usage.go:
// - praktické použití konstrukce defer
// - realizace kopie obsahu souboru
// - kratší varianta bez výpisů prováděných operací
//
// 10_11_defer_return_values.go:
// - manipulace s návratovýi hodnotami přímo v konstrukci defer
//
// 10_12_defer_arguments_evaluation.go:
// - ověření, kdy se vyhodnocují parametry předávané funkci
//   s odloženým zavoláním
// - předávání polí
//
// 10_13_multiple_defers_return_value.go:
// - manipulace s návratovýi hodnotami přímo v konstrukci defer
//
// 10_14_multiple_defers_return_value.go:
// - manipulace s návratovými hodnotami přímo v konstrukci defer
//
// 11_01_goto_loop.go:
// - nekonečná smyčka realizovaná příkazem goto
//
// 11_02_goto_end.go:
// - nekonečná smyčka realizovaná příkazem goto
// - explicitní výskok ze smyčky dalším příkazem goto
//
