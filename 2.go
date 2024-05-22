/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main

import (
	"fmt"
	"math/rand"
)


var InvSbox = []byte{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb,
	0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb,
	0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e,
	0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25,
	0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92,
	0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84,
	0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06,
	0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b,
	0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73,
	0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e,
	0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b,
	0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4,
	0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f,
	0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef,
	0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61,
	0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d}

var Sbox = []byte{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
				0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
				0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
				0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
				0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
				0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
				0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
				0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
				0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
				0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
				0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
				0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
				0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
				0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
				0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
				0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16}
// Преобразование ключа
func KEY(key string, rounds int) []string {
	keys := make([]string, rounds)
	keys[0] = key
	for i := 1; i < rounds; i++ {
		for j := 0; j < len(key); j++ {
			// Берем из Sbox
			keys[i] += string(Sbox[key[j]])
		}
	}
	return keys
}

// Добавление раундового ключа
func AddRoundKey(text, key string) string {
	for i := 0; i < len(text); i++ {
		// XOR-побит слож
		text = text[:i] + string(text[i]^key[i]) + text[i+1:]
	}
	return text
}

// Подстановка байтов
func SubBytes(text string) string {
	for i := 0; i < len(text); i++ {
		// Берем из Sbox
		text = text[:i] + string(Sbox[text[i]]) + text[i+1:]
	}
	return text
}

// Сдвиг строк влево
func ShiftRows(text string, crypt bool) string {
	retention := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// раскид по матрице
			retention[i] = append(retention[i], text[i*4+j])
		}
	}
	for i := 0; i < 4; i++ {
		storage := make([]byte, 4)
		for k := 0; k < 4; k++ {
			storage[k] = retention[i][k]
		}
		for j := 0; j < 4; j++ {
			// Для шифрования
			if crypt {
				retention[i][j] = storage[(j+i)%4]
			} else { // Для расшифрования
				retention[i][j] = storage[(j-i+4)%4]
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			text = text[:i*4+j] + string(retention[i][j]) + text[i*4+j+1:]
		}
	}
	return text
}

// Смешивание столбцов
func MixColumns(text string, initialization []int) string {
	retention := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			retention[i] = append(retention[i], text[i*4+j])
		}
	}
	initializationMat := make([][]byte, 4)
	// матрица из вектора иницилиз
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			initializationMat[i] = append(initializationMat[i], byte(initialization[(j+i)%4]))
		}
	}
	// линейн преобр
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			retention[j][i] = byte((retention[j][0] & initializationMat[j][0]) ^ (retention[j][1] & initializationMat[j][1]) ^ (retention[j][2] & initializationMat[j][2]) ^ (retention[j][3] & initializationMat[j][3]))
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			text = text[:i*4+j] + string(retention[i][j]) + text[i*4+j+1:]
		}
	}
	return text
}
func DeblockAES(text, key string, crypt bool, initialization []int) string {
	rounds := 10
	// Generate the key
	keys := KEY(key, rounds)
	fmt.Println("\nGenerated key:")
	for _, i := range keys {
		fmt.Print(i, " ")
	}
	if crypt {
		fmt.Println("\n\nEncrypted:\n")
	} else {
		fmt.Println("\nDecryption:\n")
	}
	fmt.Println("Round 1:")
	fmt.Print("AddRoundKey: ")
	// Add the round key
	retention := AddRoundKey(text, key)
	fmt.Println(retention, "\n")

	// Iterate through the rounds (starting from the second round)
	for i := 1; i < rounds-1; i++ {
		fmt.Println("Round", i+1, ":")
		// For encryption
		if crypt {
			fmt.Print("SubBytes: ")
			// Substitute bytes
			retention = SubBytes(retention)
			fmt.Println(retention)
			fmt.Print("ShiftRows: ")
			retention = ShiftRows(retention, crypt)
			fmt.Println(retention)
			fmt.Print("MixColumns: ")
			// Mix columns
			retention = MixColumns(retention, initialization)
			fmt.Println(retention)
			fmt.Print("AddRoundKey: ")
			// Add the round key
			retention = AddRoundKey(text, keys[i])
			fmt.Println(retention, "\n")
		} else { // For decryption
			fmt.Print("InvShiftRows: ")
			// Shift rows
			retention = ShiftRows(retention, crypt)
			fmt.Println(retention)
			fmt.Print("InvSubBytes: ")
			// Substitute bytes
			retention = SubBytes(retention)
			fmt.Println(retention)
			fmt.Print("AddRoundKey: ")
			retention = AddRoundKey(text, keys[i])
			fmt.Println(retention, "\n")
			fmt.Print("InvMixColumns: ")
			// Mix columns
			retention = MixColumns(retention, initialization)
			fmt.Println(retention)
		}
	}

	// Last round
	fmt.Println("Last round:")
	// For encryption
	if crypt {
		fmt.Print("SubBytes: ")
		// Substitute bytes
		retention = SubBytes(retention)
		fmt.Println(retention)
		fmt.Print("ShiftRows: ")
		// Shift rows
		retention = ShiftRows(retention, crypt)
	} else { // For decryption
		fmt.Print("InvShiftRows: ")
		// Shift rows
		retention = ShiftRows(retention, crypt)
		fmt.Println(retention)
		fmt.Print("InvSubBytes: ")
		// Substitute bytes
		retention = SubBytes(retention)
	}
	fmt.Println(retention)
	// Add the round key
	fmt.Print("AddRoundKey: ")
	retention = AddRoundKey(text, keys[rounds-1])
	fmt.Println(retention)

	fmt.Println("Final:", retention)
	return retention
}



func EncryptDecipherAES(text, key string, crypt bool, initialization []int) string {
	textTime := ""
	for i := 0; i < len(text); i += 16 {
		fmt.Println("\nBlock", i/16+1, "\n")
		retentionText := text[i : i+16]
		if !crypt {
			retentionText = DeblockAES(retentionText, key, crypt, initialization)
			key = retentionText
		} else {
			retentionKeyText := retentionText
			retentionText = DeblockAES(retentionText, key, crypt, initialization)
			key = retentionKeyText
		}
		textTime += retentionText
	}
	return textTime
}



func main() {
	text := "За свою карьеру я пропустил более 9000 бросков,проиграл почти 300 игр. 26 раз мне доверяли сделать финальный победный бросок, и я промахивался. Я терпел поражения снова, и снова, и снова. И именно поэтому я добился успеха."
	fmt.Println("Original text:", text)

	for len(text)%16 != 0 {
		text += " "
	}

	key := make([]byte, 16)
	rand.Read(key)
	fmt.Println("Initial key:", string(key))

	initialization := []int{3, 2, 4, 5}

	encryptedText := EncryptDecipherAES(text, string(key), true, initialization)
	fmt.Println("\n\nEncrypted text:", encryptedText, "\n\n")

	decryptedText := EncryptDecipherAES(encryptedText, string(key), false, initialization)
	fmt.Println("Decrypted text:", decryptedText)
}
