package main

import (
	"./HammingCodification"
	"./HuffmanCodification"
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func menuHamming() {
	var mainOp int
	r := bufio.NewReader(os.Stdin)
	continue_ := true
	for continue_ {
		clearScreen()
		fmt.Println("1 - Proteger archivo")
		fmt.Println("2 - Desporteger archivo")
		fmt.Println("3 - Introducir errores")
		fmt.Println("4 - Desproteger sin corregir errores")
		fmt.Println("5 - Volver")
		fmt.Print("Su opcion: ")
		mainOp = 0
		_, _ = fmt.Fscanf(r, "%d", &mainOp)
		switch mainOp {
		case 1:
			hamming()
		case 2:
			deHamming(true)
		case 3:
			_ = introduceErrors("")
		case 4:
			deHamming(false)
		case 5:
			continue_ = false
		}
	}
}

func menuHuffman() {
	var mainOp int
	r := bufio.NewReader(os.Stdin)
	continue_ := true
	for continue_ {
		clearScreen()
		fmt.Println("1 - Codificar")
		fmt.Println("2 - Decodificar")
		fmt.Println("3 - Volver")
		fmt.Print("Su opcion: ")
		mainOp = 0
		_, _ = fmt.Fscanf(r, "%d", &mainOp)
		_, _ = fmt.Fscanf(r, "%s")
		switch mainOp {
		case 1:
			_ = huffman("", nil)
		case 2:
			_ = desHuffman("")
		case 3:
			continue_ = false
		}
	}
}

func menuHammingHuffman() {
	var mainOp int
	r := bufio.NewReader(os.Stdin)
	continue_ := true
	for continue_ {
		clearScreen()
		fmt.Println("1 - Comprimir y proteger archivo")
		fmt.Println("2 - Desproteger y descomprimir archivo")
		fmt.Println("3 - Volver")
		fmt.Print("Su opcion: ")
		mainOp = 0
		_, _ = fmt.Fscanf(r, "%s")
		_, _ = fmt.Fscanf(r, "%d", &mainOp)
		switch mainOp {
		case 1:
			_ = preHammingHuffman("", nil)
		case 2:
			_ = preDeHammingDeHuffman("")
		case 3:
			continue_ = false
		}
	}
}

func hamming() {
	var dhOp int
	r := bufio.NewReader(os.Stdin)
	dhContinue_ := true
	for dhContinue_ {
		clearScreen()
		fmt.Println("¿Que tipo de Hamming quiere aplicar?")
		fmt.Println("1 - Hamming 7")
		fmt.Println("2 - Hamming 32")
		fmt.Println("3 - Hamming 1024")
		fmt.Println("4 - Hamming 32768")
		fmt.Println("5 - Volver")
		fmt.Printf("Su opcion: ")
		dhOp = 0
		_, _ = fmt.Fscanf(r, "%d", &dhOp)
		switch dhOp {
		case 1:
			_ = preHamming(7, "", nil)
		case 2:
			_ = preHamming(32, "", nil)
		case 3:
			_ = preHamming(1024, "", nil)
		case 4:
			_ = preHamming(32768, "", nil)
		case 5:
			dhContinue_ = false
		}
	}
}

func deHamming(fixErrors bool) {
	var dhOp int
	r := bufio.NewReader(os.Stdin)
	dhContinue_ := true
	for dhContinue_ {
		clearScreen()
		fmt.Println("¿Que tipo de Hamming ha sido aplicado?")
		fmt.Println("1 - Hamming 7")
		fmt.Println("2 - Hamming 32")
		fmt.Println("3 - Hamming 1024")
		fmt.Println("4 - Hamming 32768")
		fmt.Println("5 - Volver")
		fmt.Printf("Su opcion: ")
		dhOp = 0
		_, _ = fmt.Fscanf(r, "%d", &dhOp)
		switch dhOp {
		case 1:
			_ = preDeHamming(7, "", fixErrors)
		case 2:
			_ = preDeHamming(32, "", fixErrors)
		case 3:
			_ = preDeHamming(1024, "", fixErrors)
		case 4:
			_ = preDeHamming(32768, "", fixErrors)
		case 5:
			dhContinue_ = false
		}
	}
}

func introduceErrors(fileName string) error {
	var body, fileWithErrors []byte
	var err error
	originalText, err := loadFile(fileName, false)
	if err != nil {
		return err
	}
	//Split the string between name and extension
	extension := strings.Split(fileName, ".")
	switch extension[1] {
	case "ha1":
		body = originalText[:len(originalText)-10]
		fileWithErrors = append(HammingCodification.InsertError7(body), originalText[len(originalText)-10:]...)
	case "ha2":
		body = originalText[:len(originalText)-20]
		fileWithErrors = append(HammingCodification.InsertError(body, 32), originalText[len(originalText)-20:]...)
	case "ha3":
		body = originalText[:len(originalText)-20]
		fileWithErrors = append(HammingCodification.InsertError(body, 1024), originalText[len(originalText)-20:]...)
	case "ha4":
		body = originalText[:len(originalText)-20]
		fileWithErrors = append(HammingCodification.InsertError(body, 32768), originalText[len(originalText)-20:]...)
	default:
		return errors.New("La extension del archivo no es válida.")
	}
	_ = saveFile(strings.Replace(fileName, ".ha", ".he", -1), fileWithErrors)
	return nil
}

//preHamming size type of Hamming, unixDate date since when you can decode the file
func preHamming(size int, fileName string, unixDate []byte) error {
	var encodedBody []byte
	body, err := loadFile(fileName, false)
	var fileType string
	if err != nil {
		return err
	} else {
		switch size {
		case 7:
			fileType = ".ha1"
		case 32:
			fileType = ".ha2"
		case 1024:
			fileType = ".ha3"
		case 32768:
			fileType = ".ha4"
		}
		if len(body) == 0 {
			encodedBody = []byte{}
			if size != 7 {
				for i := 0; i < 10; i++ {
					encodedBody = append(encodedBody, byte(48))
				}
			}
		} else {
			switch size {
			case 7:
				encodedBody = HammingCodification.Hamming7(body)
			case 32:
				encodedBody = HammingCodification.Hamming(size, body)
			case 1024:
				encodedBody = HammingCodification.Hamming(size, body)
			case 32768:
				encodedBody = HammingCodification.Hamming(size, body)
			}
		}
		fileName = strings.Replace(fileName, ".txt", fileType, -1)
		encodedBody = append(encodedBody, unixDate...)
		err = saveFile(fileName, encodedBody)
		if err != nil {
			return err
		}

	}
	return nil
}

func preDeHamming(size int, fileName string, fixErrors bool) error {
	var body []byte
	var err error
	var hammingCase string
	extension := strings.Split(fileName, ".")
	switch extension[1] {
	case "ha1":
		hammingCase = "1"
	case "ha2":
		hammingCase = "2"
	case "ha3":
		hammingCase = "3"
	case "ha4":
		hammingCase = "4"
	case "he1":
		hammingCase = "1"
	case "he2":
		hammingCase = "2"
	case "he3":
		hammingCase = "3"
	case "he4":
		hammingCase = "4"
	default:
		return errors.New("La extension del archivo no es válida.")
	}

	body, err = loadFile(fileName, true)
	if err != nil {
		return err
	}
	var decodedFile []byte
	if len(body) == 0 {
		decodedFile = []byte{}
	} else {
		if size == 7 {
			decodedFile = HammingCodification.DeHamming7(body, fixErrors)
		} else {
			decodedFile = HammingCodification.CallDecode(size, body, fixErrors)
		}
	}
	if fixErrors {
		fileName = strings.Replace(fileName, "."+extension[1], ".dh"+hammingCase, -1)
	} else {
		fileName = strings.Replace(fileName, "."+extension[1], ".de"+hammingCase, -1)
	}
	err = saveFile(fileName, decodedFile)
	if err != nil {
		return err
	}
	return nil
}

func statistics(url string) []string {
	extension := strings.Split(url, ".")[1]
	urlWithoutExtension := strings.Split(url, ".")[0]
	extensions := []string{"." + extension, ".ha1", ".ha2", ".ha3", ".ha4", ".huf", ".dic", ".hh1", ".dichh1", ".hh2", ".dichh2", ".hh3", ".dichh3", ".hh4", ".dichh4"}
	ret := make([]string, 0)
	for index := 0; index < len(extensions); index++ {
		body, err := loadFile(urlWithoutExtension+extensions[index], false)
		if err == nil {
			switch extensions[index] {
			case "." + extension:
				ret = append(ret, "El archivo original tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".ha1":
				ret = append(ret, "Hamming 7 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".ha2":
				ret = append(ret, "Hamming 32 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".ha3":
				ret = append(ret, "Hamming 1024 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".ha4":
				ret = append(ret, "Hamming 32768 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".huf":
				ret = append(ret, "Huffman tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".dic":
				ret = append(ret, "La tabla de Huffman tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".hh1":
				ret = append(ret, "Hamming/Huffman 7 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".dichh1":
				ret = append(ret, "La tabla de Hamming/Huffman 7 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".hh2":
				ret = append(ret, "Hamming/Huffman 32 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".dichh2":
				ret = append(ret, "La tabla de Hamming/Huffman 32 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".hh3":
				ret = append(ret, "Hamming/Huffman 1024 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".dichh3":
				ret = append(ret, "La tabla de Hamming/Huffman 1024 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".hh4":
				ret = append(ret, "Hamming/Huffman 32768 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			case ".dichh4":
				ret = append(ret, "La tabla de Hamming/Huffman 32768 tiene un tamaño de: "+strconv.Itoa(len(body))+" Bytes "+" o "+strconv.Itoa((len(body))/1024)+" KB")
			}
		} else {
			ret = append(ret, "" /*"No se encontro el archivo con extension "+extensions[index]+"."*/)
		}
	}
	return ret
}

func huffman(fileName string, unixDate []byte) error {
	//Since golang does not show the time a program runs...
	body, err := loadFile(fileName, false)
	if err != nil {
		return err
	} else {
		encodedBody, dictionary := HuffmanCodification.CallHuffman(body)
		dictionary = append(dictionary, unixDate...)
		fileName = strings.Split(fileName, ".")[0]
		fileName = fileName + ".huf"
		err = saveFile(fileName, encodedBody)
		if err != nil {
			return err
		}
		fileName = strings.Replace(fileName, "huf", "dic", -1)
		err = saveFile(fileName, dictionary)
		if err != nil {
			return err
		}
	}
	return nil
}

func desHuffman(fileName string) error {
	body, err := loadFile(fileName+".huf", false)
	if err != nil {
		return err
	}
	table, err := loadFile(fileName+".dic", true)
	if err != nil {
		return err
	} else {
		decodedBody := HuffmanCodification.Deshuffman(body, table)
		fileName = fileName + ".dhu"
		err = saveFile(fileName, decodedBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func preHammingHuffman(fileName string, unixDate []byte) error {
	//Ask for hamming type
	var dhOp int
	dhContinue_ := true
	var size int
	var bodyExtension string
	var dicExtension string
	for dhContinue_ {
		clearScreen()
		fmt.Println("¿Que tipo de Hamming quiere aplicar?")
		fmt.Println("1 - Hamming 7")
		fmt.Println("2 - Hamming 32")
		fmt.Println("3 - Hamming 1024")
		fmt.Println("4 - Hamming 32768")
		fmt.Println("5 - Volver")
		fmt.Printf("Su opcion: ")
		dhOp = 0
		switch dhOp {
		case 1:
			size = 7
			bodyExtension = ".hh1"
			dicExtension = ".dichh1"
			dhContinue_ = false
		case 2:
			size = 32
			bodyExtension = ".hh2"
			dicExtension = ".dichh2"
			dhContinue_ = false
		case 3:
			size = 1024
			bodyExtension = ".hh3"
			dicExtension = ".dichh3"
			dhContinue_ = false
		case 4:
			size = 32768
			bodyExtension = ".hh4"
			dicExtension = ".dichh4"
			dhContinue_ = false
		case 5:
			return nil
		}
	}
	//Ask for date

	//Ask for file name
	var body []byte
	err := errors.New("Not nil error")
	for err != nil {
		body, err = loadFile(fileName, false)
		if err != nil {
			return nil
		}

	}

	//Compress
	compressedBody, dictionary := HuffmanCodification.CallHuffman(body)

	//Protect
	var encodedBody []byte
	var encodedDic []byte
	if len(compressedBody) == 0 {
		encodedBody = []byte{}
		if size != 7 {
			for i := 0; i < 10; i++ {
				encodedBody = append(encodedBody, byte(48))
				encodedDic = append(encodedDic, byte(48))
			}
		}
	} else {
		switch size {
		case 7:
			encodedBody = HammingCodification.Hamming7(compressedBody)
			encodedDic = HammingCodification.Hamming7(dictionary)
		default:
			encodedBody = HammingCodification.Hamming(size, compressedBody)
			encodedDic = HammingCodification.Hamming(size, dictionary)
		}
	}
	encodedDic = append(encodedDic, unixDate...)

	//Save files
	fileName = strings.Split(fileName, ".")[0]
	err = saveFile(fileName+bodyExtension, encodedBody)
	if err != nil {
		return err
	}
	err = saveFile(fileName+dicExtension, encodedDic)
	if err != nil {
		return err
	}
	return nil
}

func preDeHammingDeHuffman(fileName string) error {

	//Ask for hamming type
	var dhOp int
	r := bufio.NewReader(os.Stdin)
	dhContinue_ := true
	var size int
	var bodyExtension string
	var dicExtension string
	var finalExtension string
	extension := strings.Split(fileName, ".")
	for dhContinue_ {
		fmt.Println("¿Que tipo de Hamming ha sido aplicado?")
		fmt.Println("1 - Hamming 7")
		fmt.Println("2 - Hamming 32")
		fmt.Println("3 - Hamming 1024")
		fmt.Println("4 - Hamming 32768")
		fmt.Println("5 - Volver")
		fmt.Printf("Su opcion: ")
		dhOp = 0
		_, _ = fmt.Fscanf(r, "%d", &dhOp)
		switch extension[1] {
		case ".hh1":
			size = 7
			bodyExtension = ".hh1"
			dicExtension = ".dichh1"
			finalExtension = ".dhh1"
			dhContinue_ = false
		case ".hh2":
			size = 32
			bodyExtension = ".hh2"
			dicExtension = ".dichh2"
			finalExtension = ".dhh2"
			dhContinue_ = false
		case ".hh3":
			size = 1024
			bodyExtension = ".hh3"
			dicExtension = ".dichh3"
			finalExtension = ".dhh3"
			dhContinue_ = false
		case ".hh4":
			size = 32768
			bodyExtension = ".hh4"
			dicExtension = ".dichh4"
			finalExtension = ".dhh4"
			dhContinue_ = false
		default:
			return errors.New("Extension invalida.")
		}
	}

	//Ask for file name

	var encodedBody []byte
	var encodedDic []byte
	err1 := errors.New("Not nil error")
	err2 := errors.New("Not nil error")
	fileName = extension[0]
	encodedBody, err1 = loadFile(fileName+bodyExtension, false)
	if err1 != nil {
		return err1
	}
	encodedDic, err2 = loadFile(fileName+dicExtension, true)
	if err2 != nil {
		return err2
	}

	//Desprotect
	var decodedBody []byte
	var decodedDic []byte
	if size == 7 {
		decodedBody = HammingCodification.DeHamming7(encodedBody, true)
		decodedDic = HammingCodification.DeHamming7(encodedDic, true)
	} else {
		decodedBody = HammingCodification.CallDecode(size, encodedBody, true)
		decodedDic = HammingCodification.CallDecode(size, encodedDic, true)
	}

	//Descompress
	descompressBody := HuffmanCodification.Deshuffman(decodedBody, decodedDic)

	//Save file
	err := saveFile(fileName+finalExtension, descompressBody)
	if err != nil {
		return err
	}
	return nil

}

/*func askDate() ([]byte, error) {
	//Ask for the date
	clearScreen()
	r := bufio.NewReader(os.Stdin)
	var auxDay, auxMonth, auxYear, auxHour, auxMinutes, auxSeconds string
	var day, month, year, hour, minutes, seconds int
	err := make([]error, 6)
	fmt.Println("Ingrese el dia, mes, año, hora, minutos y segundos en los que quiere la decodificacion del archivo este disponible:")
	fmt.Print("Dia: ")
	_, _ = fmt.Fscanf(r, "%s", &auxDay)
	_, _ = fmt.Fscanf(r, "%d")
	fmt.Print("Mes: ")
	_, _ = fmt.Fscanf(r, "%s", &auxMonth)
	_, _ = fmt.Fscanf(r, "%d")
	fmt.Print("Año: ")
	_, _ = fmt.Fscanf(r, "%s", &auxYear)
	_, _ = fmt.Fscanf(r, "%d")
	fmt.Print("Hora: ")
	_, _ = fmt.Fscanf(r, "%s", &auxHour)
	_, _ = fmt.Fscanf(r, "%d")
	fmt.Print("Minutos: ")
	_, _ = fmt.Fscanf(r, "%s", &auxMinutes)
	_, _ = fmt.Fscanf(r, "%d")
	fmt.Print("Segundos: ")
	_, _ = fmt.Fscanf(r, "%s", &auxSeconds)
	_, _ = fmt.Fscanf(r, "%d")
	//Searching errors process
	day, err[0] = strconv.Atoi(auxDay)
	month, err[1] = strconv.Atoi(auxMonth)
	year, err[2] = strconv.Atoi(auxYear)
	hour, err[3] = strconv.Atoi(auxHour)
	minutes, err[4] = strconv.Atoi(auxMinutes)
	seconds, err[5] = strconv.Atoi(auxSeconds)
	//Check if the date have errors
	for i := 0; i < len(err); i++ {
		if err[i] != nil {
			return nil, errors.New("Formato de fecha incorrecto.")
		}
	}
	//No error found then create the date
	parseMonth := time.Month(month)
	location, _ := time.LoadLocation("America/Argentina/Cordoba")
	auxDate := time.Date(year, parseMonth, day, hour, minutes, seconds, 0, location)
	auxUnixDate := auxDate.Unix()
	s := []byte(strconv.FormatInt(auxUnixDate, 10))
	unixDate := []byte(s)
	for i := len(unixDate); i < 10; i = len(unixDate) {
		unixDate = append([]byte{48}, unixDate...)
	}
	clearScreen()
	return unixDate, nil
}*/

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	fmt.Println("#####################################")
	fmt.Println("___________Hamming/Huffman___________")
	fmt.Println("#####################################")
	fmt.Println()
}
