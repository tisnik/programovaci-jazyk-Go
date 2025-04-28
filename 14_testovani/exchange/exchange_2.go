// získání a zpracování kurzovního lístku
func get_exchange_rate(code string) float64 {
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
