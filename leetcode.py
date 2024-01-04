# def roman_to_integer(roman):
#     roman_dict = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
#     result = 0

#     prev_value = 0
#     for char in roman:
#         value = roman_dict[char]

#         # Jika nilai saat ini lebih besar dari nilai sebelumnya, tambahkan ke hasil
#         if value > prev_value:
#             result += value - 2 * prev_value
#         else:
#             result += value

#         prev_value = value

#     return result

# # Contoh penggunaan:
# roman_numeral = input("Masukan angka romawi: ")
# integer_value = roman_to_integer(roman_numeral)
# print(f"{roman_numeral} dalam angka biasa adalah {integer_value}")

bahasa_romawi = {
    "I" or "i" : 1,
    "V" or "v" : 5,
    "X" or "x" : 10,
    "L" or "l" : 50,
    "C" or "c" : 100,
    "D" or "d" : 500,
    "M" or "m" : 1000
}

bil1 = int(input("Bil1: "))
bil2 = int(input("Bil2: "))
hasil = int(bil1 / bil2)
print(hasil)
a = ["Nol", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh"]
# print (a[hasil])

hasil1 = bil1 % bil2
print (a[hasil] + " " + a[hasil1])