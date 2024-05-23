package main

import (
	"fmt"
	"math/rand"
	"time"
)

var S_box [256]byte
var InvS_box [256]byte

func createSbox(S_box *[256]byte) {
	a := 75
	for i := 0; i < 256; i++ {
		S_box[i] = byte(a)
		a += 213
		a %= 256
	}
}

func createInvSbox(S_box *[256]byte, InvS_box *[256]byte) {
	for i := 0; i < 256; i++ {
		InvS_box[S_box[i]] = byte(i)
	}
}

func keyExpansion(key string, S_box *[256]byte, rounds int) []string {
	keys := make([]string, rounds)
	keys[0] = key
	for i := 1; i < rounds; i++ {
		for j := 0; j < len(key); j++ {
			keys[i] += string(S_box[keys[i-1][j]])
		}
	}
	return keys
}

func AddRoundKey(text, roundKey string) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ roundKey[i]
	}
	return string(result)
}

func SubBytes(text string, S_box *[256]byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = S_box[text[i]]
	}
	return string(result)
}

func ShiftRows(text string) string {
	mat := [4][4]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mat[i][j] = text[i*4+j]
		}
	}
	for i := 0; i < 4; i++ {
		t := [4]byte{}
		for k := 0; k < 4; k++ {
			t[k] = mat[i][k]
		}
		for j := 0; j < 4; j++ {
			mat[i][j] = t[(j+i)%4]
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			text = text[:i*4+j] + string(mat[i][j]) + text[i*4+j+1:]
		}
	}
	return text
}

func InvShiftRows(text string) string {
	mat := [4][4]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mat[i][j] = text[i*4+j]
		}
	}
	for i := 0; i < 4; i++ {
		t := [4]byte{}
		for k := 0; k < 4; k++ {
			t[k] = mat[i][k]
		}
		for j := 0; j < 4; j++ {
			mat[i][j] = t[(j-i+4)%4]
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			text = text[:i*4+j] + string(mat[i][j]) + text[i*4+j+1:]
		}
	}
	return text
}

func MixColumns(text string, iniz []int) string {
	mat := [4][4]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mat[i][j] = text[i*4+j]
		}
	}
	vecMat := [4][4]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			vecMat[i][(j+i)%4] = byte(iniz[j])
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mat[j][i] = (mat[j][0] & vecMat[j][0]) ^ (mat[j][1] & vecMat[j][1]) ^ (mat[j][2] & vecMat[j][2]) ^ (mat[j][3] & vecMat[j][3])
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			text = text[:i*4+j] + string(mat[i][j]) + text[i*4+j+1:]
		}
	}
	return text
}

func blockAES(text, key string, S_box *[256]byte, iniz []int) string {
	rounds := 10
	keys := keyExpansion(key, S_box, rounds)

	fmt.Println("\nGenerated keys: ")
	for _, k := range keys {
		fmt.Print(k, " ")
	}
	fmt.Println("\nEncryption:\n")
	fmt.Println("round 1:")
	fmt.Print("AddRoundKey: ")
	tempEnc := AddRoundKey(text, key)
	fmt.Println(tempEnc, "\n")
	for i := 1; i < rounds-1; i++ {
		fmt.Println("round", i+1, ":")
		fmt.Print("SubBytes: ")
		tempEnc = SubBytes(tempEnc, S_box)
		fmt.Println(tempEnc)
		fmt.Print("ShiftRows: ")
		tempEnc = ShiftRows(tempEnc)
		fmt.Println(tempEnc)
		fmt.Print("MixColumns: ")
		tempEnc = MixColumns(tempEnc, iniz)
		fmt.Println(tempEnc)
		fmt.Print("AddRoundKey: ")
		tempEnc = AddRoundKey(text, keys[i])
		fmt.Println(tempEnc, "\n")
	}
	fmt.Println("last round:")
	fmt.Print("SubBytes: ")
	tempEnc = SubBytes(tempEnc, S_box)
	fmt.Println(tempEnc)
	fmt.Print("ShiftRows: ")
	tempEnc = ShiftRows(tempEnc)
	fmt.Println(tempEnc)
	fmt.Print("AddRoundKey: ")
	tempEnc = AddRoundKey(text, keys[rounds-1])
	fmt.Println(tempEnc)
	fmt.Println("final cypher:", tempEnc)

	return tempEnc
}

func deBlockAES(text, key string, InvS_Box, S_box *[256]byte, invIniz []int) string {
	rounds := 10
	keys := keyExpansion(key, S_box, rounds)
	fmt.Println("Decryption:\n")
	fmt.Println("round 1:")
	tempDec := AddRoundKey(text, key)
	fmt.Println(tempDec, "\n")
	for i := 1; i < rounds-1; i++ {
		fmt.Println("round", i+1, ":")
		fmt.Print("InvSubBytes: ")
		tempDec = SubBytes(tempDec, InvS_Box)
		fmt.Println(tempDec)
		fmt.Print("InvShiftRows: ")
		tempDec = InvShiftRows(tempDec)
		fmt.Println(tempDec)
		fmt.Print("InvMixColumns: ")
		tempDec = MixColumns(tempDec, invIniz)
		fmt.Println(tempDec)
		fmt.Print("AddRoundKey: ")
		tempDec = AddRoundKey(text, keys[i])
		fmt.Println(tempDec, "\n")
	}
	fmt.Println("last round:")
	fmt.Print("InvSubBytes: ")
	tempDec = SubBytes(tempDec, InvS_Box)
	fmt.Println(tempDec)
	fmt.Print("InvShiftRows: ")
	tempDec = InvShiftRows(tempDec)
	fmt.Println(tempDec)
	fmt.Print("AddRoundKey: ")
	tempDec = AddRoundKey(text, keys[rounds-1])
	fmt.Println(tempDec)
	fmt.Println("final text:", tempDec)
	return tempDec
}

func AESEnc(text, key string, S_box *[256]byte, iniz []int) string {
	cypher := ""
	for i := 0; i < len(text); i += 16 {
		fmt.Println("\nblock number", i/16+1, "\n")
		tempText := text[i : i+16]
		tempkey := tempText
		tempText = blockAES(tempText, key, S_box, iniz)
		key = tempkey
		cypher += tempText
	}
	return cypher
}

func AESDec(cypher, key string, S_box, InvS_box *[256]byte, iniz []int) string {
	text := ""
	for i := 0; i < len(cypher); i += 16 {
		fmt.Println("\nblock number", i/16+1, "\n")
		tempText := cypher[i : i+16]
		tempText = deBlockAES(tempText, key, InvS_box, S_box, iniz)
		key = tempText
		text += tempText
	}
	return text
}

func main() {
	text := "все конец"

	key := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		key += string(rand.Intn(256))
	}
	fmt.Println("key:", key, "\n")

	createSbox(&S_box)
	createInvSbox(&S_box, &InvS_box)

	for i := 0; i < 256; i++ {
		fmt.Print(S_box[i], ", ")
	}
	fmt.Println()
	for i := 0; i < 256; i++ {
		fmt.Print(InvS_box[i], ", ")
	}
	fmt.Println()

	iniz := []int{2, 3, 1, 1}

	k := 0
	for len(text)%16 != 0 {
		text += " "
		k++
	}

	text = AESEnc(text, key, &S_box, iniz)

	fmt.Println("\n\nCYPHER:\n")
	fmt.Println(text, "\n")

	text = AESDec(text, key, &S_box, &InvS_box, iniz)

	for k != 0 {
		text = text[:len(text)-1]
		k--
	}

	fmt.Println(text)
}


