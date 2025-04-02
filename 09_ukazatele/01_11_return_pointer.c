#include <stdio.h>

int* get_pointer(void) {
    // lokální proměnná
    int i = 42;

    // vracíme adresu lokální proměnné
    return &i;
}

int main(void) {
    // získání adresy lokální proměnné
    int *p = get_pointer();

    // tisk adresy
    printf("%p\n", p);

    // přístup na adresu
    printf("%d\n", *p);

    return 0;
}
