// převod měny na základě reálného kurzovního lístku
func get_exchange_rate_from_url(code string) float64 {
	const URL = "https://www.cnb.cz/cs/financni_trhy/devizovy_trh/kurzy_devizoveho_trhu/denni_kurz.txt"

	response, err := http.Get(URL)
	if err != nil {
		panic("Connection refused")
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			code_str := s[3]
			rate_str := strings.Replace(s[4], ",", ".", 1)
			if code == code_str {
				rate_f, err := strconv.ParseFloat(rate_str, 64)
				if err != nil {
					panic(err)
				}
				return rate_f
			}
		}
	}

	return 0
}

// převod měny na základě kurzovního lístku přečteného ze souboru
func get_exchange_rate_from_file(code string) float64 {
	const FILENAME = "kurzy.txt"

	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		if len(s) == 5 {
			code_str := s[3]
			rate_str := strings.Replace(s[4], ",", ".", 1)
			if code == code_str {
				rate_f, err := strconv.ParseFloat(rate_str, 64)
				if err != nil {
					panic(err)
				}
				return rate_f
			}
		}
	}

	return 0
}

// pouze napodobení chování funkce pro převod měny
func mocked_get_exchange_rate(code string) float64 {
	return 21.5
}
