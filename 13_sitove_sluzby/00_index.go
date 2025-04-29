// Síťové služby
//
// Seznam příkladů uložených v tomto podadresáři:
//
// 01_01_simple_client.go:
// - standardní balíček net
// - síťový klient, který přečte ze serveru sekvenci bajtů
//
// 01_02_simple_client_headers.go:
// - standardní balíček net
// - upravený klient, který vytiskne místní i vzdálenou adresu
//
// 01_03_simple_server.go:
// - standardní balíček net
// - server naslouchající na portu 1234
// - omezení na místní připojení
// - po navázání připojení se odešle jeden bajt a připojení se ukončí
//
// 01_04_simple_server_no_localhost.go:
// - standardní balíček net
// - server naslouchající na portu 1234
// - po navázání připojení se odešle jeden bajt a připojení se ukončí
//
// 01_05_slow_server.go:
// - standardní balíček net
// - server odpovídající klientovi opožděně
// - omezení na místní připojení
// - po navázání připojení se odešle jeden bajt a připojení se ukončí
//
// 01_06_multi_connection_server.go:
// - standardní balíček net
// - server, který dokáže obsloužit více klientů současně
// - po navázání připojení se odešle jeden bajt a připojení se ukončí
//
// 01_07_text_server.go:
// - standardní balíček net
// - server posílající klientovi textová data
//
// 01_08_better_text_server.go:
// - standardní balíček net
// - server posílající klientovi textová data
// - uzavření připojení v bloku defer
//
// 01_09_text_client.go:
// - standardní balíček net
// - jednoduchý klient akceptující celý textový řádek
//
// 01_10_lookup.go:
// - standardní balíček net
// - překlad doménového jména na IP adresy
//
// 01_11_parse_ip.go:
// - standardní balíček net
// - parsing IPv4 a IPv6 adresy
//
// 01_12_ipv4_constructor.go:
// - standardní balíček net
// - konstruktor adresy typu IPv4
//
// 02_01_http_get.go:
// - standardní balíček net/http
// - použití HTTP metody GET
//
// 02_02_http_print_headers.go:
// - standardní balíček net/http
// - vytištění hlavičky HTTP odpovědi
//
// 02_03_http_server.go:
// - standardní balíček net/http
// - nejjednodušší HTTP server s jediným endpointem
//
// 02_04_http_server_with_state.go:
// - standardní balíček net/http
// - HTTP server se dvěma endpointy
// - nekorektní práce se stavovou proměnnou
//
// 02_05_http_server_with_state_mutex.go:
// - standardní balíček net/http
// - HTTP server se dvěma endpointy
// - korektní práce se stavovou proměnnou
// - synchronizace přístupu přes mutex
//
// 02_06_file_server.go:
// - standardní balíček net/http
// - HTTP server vracející statický obsah
//
// 02_07_custom_server.go:
// - standardní balíček net/http
// - kombinace předchozích možností HTTP serverů
//
// 02_08_http_server_write_bytes.go:
// - standardní balíček net/http
// - HTTP server posílající sekvenci bajtů
//

