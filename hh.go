package main

import (
	"./HammingCodification"
	"./HuffmanCodification"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func introduceErrors(fileName string) error {
	var body, fileWithErrors []byte
	var err error
	originalText, err := loadFile(fileName, false)
	if err != nil {
		return err
	}
	//Split the string between name and extension
	extension := strings.Split(fileName, ".")
	switch extension[len(extension)-1] {
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
			fileType = "ha1"
		case 32:
			fileType = "ha2"
		case 1024:
			fileType = "ha3"
		case 32768:
			fileType = "ha4"
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
		extension := strings.Split(fileName, ".")
		fileName = strings.Replace(fileName, extension[len(extension)-1], fileType, -1)
		encodedBody = append(encodedBody, unixDate...)
		err = saveFile(fileName, encodedBody)
		if err != nil {
			return err
		}

	}
	return nil
}

func preDeHamming(fileName string, fixErrors bool) error {
	var body []byte
	var err error
	var hammingCase string
	var size int
	extension := strings.Split(fileName, ".")
	switch extension[len(extension)-1] {
	case "ha1":
		hammingCase = "1"
		size = 7
	case "ha2":
		hammingCase = "2"
		size = 32
	case "ha3":
		hammingCase = "3"
		size = 1024
	case "ha4":
		hammingCase = "4"
		size = 32768
	case "he1":
		hammingCase = "1"
		size = 7
	case "he2":
		hammingCase = "2"
		size = 32
	case "he3":
		hammingCase = "3"
		size = 1024
	case "he4":
		hammingCase = "4"
		size = 32768
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
		fileName = strings.Replace(fileName, "."+extension[len(extension)-1], ".dh"+hammingCase, -1)
	} else {
		fileName = strings.Replace(fileName, "."+extension[len(extension)-1], ".de"+hammingCase, -1)
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
	body, err := loadFile(fileName, false)
	if err != nil {
		return err
	}
	fileName = strings.Replace(fileName, "huf", "dic", -1)
	table, err := loadFile(fileName, true)
	if err != nil {
		return err
	} else {
		decodedBody := HuffmanCodification.Deshuffman(body, table)
		fileName = strings.Replace(fileName, "dic", "dhu", -1)
		err = saveFile(fileName, decodedBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func preHammingHuffman(size int, fileName string, unixDate []byte) error {
	//Ask for hamming type
	var bodyExtension string
	var dicExtension string
	switch size {
	case 7:
		bodyExtension = ".hh1"
		dicExtension = ".dichh1"
	case 32:
		bodyExtension = ".hh2"
		dicExtension = ".dichh2"
	case 1024:
		bodyExtension = ".hh3"
		dicExtension = ".dichh3"
	case 32768:
		bodyExtension = ".hh4"
		dicExtension = ".dichh4"
	}

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
	var finalExtension string
	var dicExtension string
	var size int
	extension := strings.Split(fileName, ".")
	switch extension[len(extension)-1] {
	case "hh1":
		dicExtension = ".dichh1"
		finalExtension = ".dhh1"
		size = 7
	case "hh2":
		dicExtension = ".dichh2"
		finalExtension = ".dhh2"
		size = 32
	case "hh3":
		dicExtension = ".dichh3"
		finalExtension = ".dhh3"
		size = 1024
	case "hh4":
		dicExtension = ".dichh4"
		finalExtension = ".dhh4"
		size = 32768
	default:
		return errors.New("La extension del archivo no es válida.")
	}

	var encodedBody []byte
	var encodedDic []byte
	err1 := errors.New("Not nil error")
	err2 := errors.New("Not nil error")

	encodedBody, err1 = loadFile(fileName, false)
	if err1 != nil {
		return err1
	}

	encodedDic, err2 = loadFile(strings.Split(fileName, ".")[0]+dicExtension, true)
	if err2 != nil {
		return err2
	}

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
	err := saveFile(strings.Split(fileName, ".")[0]+finalExtension, descompressBody)
	if err != nil {
		return err
	}

	return nil
}
