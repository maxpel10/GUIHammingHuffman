package HammingCodification

import (
	"fmt"
	"math"
)

// This function call to "takeBits" until the bits array is empty.
func callTakeBits(hMBits int, body []byte) (Array [][]byte) {
	var TrashBits int
	var bytes []byte

	TrashBits = 0

	for body != nil {
		bytes, body, TrashBits = takeBits(hMBits, body, TrashBits)

		Array = append(Array, bytes)
	}
	return Array

}

// This function take the amount the bits you want for a bytes array. ( must be more than 7 )
func takeBits(bits int, body []byte, NumberOfTrashBits int) ([]byte, []byte, int) {

	var cantByte int
	var cantBit int
	var arrBit []byte
	var mask uint8
	var bait byte
	var finish bool
	var bitsToMove int

	// initialise variables

	cantByte = bits / 8                // amount bytes i need.
	bitsToMove = 8 - NumberOfTrashBits // how many bits i need to shift.

	cantBit = bits - (cantByte * 8) // amount bits i need for the incomplete byte.

	if bits >= 8 {
		if len(body) <= cantByte {
			finish = true
		}
		for index := 0; index <= cantByte; index++ {
			if len(body) <= cantByte+1 {
				body = append(body, uint8(0))
			}

			// adjust

			baitAux := body[index]
			baitAux = baitAux << uint(NumberOfTrashBits) // shift to left how many bits i need to remove
			nextBait := body[index+1]
			nextBait = nextBait >> uint(bitsToMove)
			aux := baitAux | nextBait // merge the bytes to pass the bits from the next byte to this one.

			arrBit = append(arrBit, aux) // put it on the array.

		}

		if finish == true {
			if cantBit > 0 {
				bait = arrBit[cantByte] // adjust the byte

				mask = doMask(cantBit) // make the mask by how many bits i need
				bait = bait & mask     // make the byte

				arrBit[cantByte] = bait // put the byte on the array.

			}
			body = nil
		} else {
			if cantBit > 0 {
				bait = arrBit[cantByte] // adjust the byte
				mask = doMask(cantBit)  // make the mask by how many bits i need
				bait = bait & mask      // make the byte
				arrBit[cantByte] = bait // put the byte on the array.

			}
			NumberOfTrashBits += cantBit

			if NumberOfTrashBits >= 8 {

				cutPosition := cantByte + (NumberOfTrashBits / 8)
				NumberOfTrashBits = NumberOfTrashBits - 8
				body = body[cutPosition:]
			} else {
				body = body[cantByte:] // adjust the array
			}

		}
		return arrBit, body, NumberOfTrashBits
	} else {
		fmt.Printf("This function is not available for values less than 8 bits.")
		return nil, nil, 0
	}
}

// This function make a mask to take bits from a byte (left to right).
func doMask(bits int) uint8 {
	if bits > 8 {
		return uint8(0)
	} else if bits < 0 {
		return uint8(0)
	} else {
		valMask := math.Pow(2, float64(bits)) - 1
		mask := uint8(valMask) << uint(8-bits)
		return mask
	}

}
