package src

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"net"
)

type TConnection struct {
	TSelRemote             []byte
	TSelLocal              []byte
	Socket                 *net.Conn
	maxTPduSizeParam       int
	messageTimeout         int
	messageFragmentTimeout int
	serverThread           *ServerThread
	Closed                 bool
	os                     *bufio.Writer
	is                     *bufio.Reader
	srcRef                 int
	dstRef                 int
	maxTPduSize            int
}

func NewTConnection(socket *net.Conn, maxTPduSizeParam int, messageTimeout int, messageFragmentTimeout int, serverThread *ServerThread) *TConnection {
	if maxTPduSizeParam < 7 || maxTPduSizeParam > 16 {
		throw("maxTPduSizeParam is incorrect")
	}

	reader := bufio.NewReaderSize(*socket, 8192)
	writeByteByter := bufio.NewWriter(*socket)
	var maxTPduSize int
	if maxTPduSizeParam == 16 {
		maxTPduSize = 65531
	} else {

		maxTPduSize = int(math.Pow(2, float64(maxTPduSizeParam)))
	}

	return &TConnection{
		Socket:                 socket,
		maxTPduSizeParam:       maxTPduSizeParam,
		messageTimeout:         messageTimeout,
		messageFragmentTimeout: messageFragmentTimeout,
		serverThread:           serverThread,
		Closed:                 false,
		os:                     writeByteByter,
		is:                     reader,
		maxTPduSize:            maxTPduSize,
	}
}

/**
 * Listens for a NewTPDU and writeByteBytes the extracted TSDU into the passed buffer.
 *
 * @param tSduBuffer the buffer that is filled with the received TSDU data.
 * @throws EOFException if a Disconnect Request (DR) was received or the socket was simply Closed
 * @throws SocketTimeoutException if a messageFragmentTimeout is thrown by the socket while
 *     receiving the remainder of a message
 * @throws IOException if an ErrorPDU (ER) was received, any syntax error in the received message
 *     header was detected or the tSduBuffer is too small to hold the complete PDU.
 * @throws TimeoutException this exception is thrown if the first byte of Newmessage is not
 *     received within the message timeout.
 */
func (t *TConnection) receive(tSduBuffer *bytes.Buffer) {
	//socket := *t.Socket
	is := t.is

	packetLength := 0
	eot := 0
	li := 0
	tPduCode := 0
	//duration, err := time.ParseDuration(fmt.Sprintf("%dms", t.messageTimeout))
	//if err != nil {
	//	panic(err)
	//}
	//err = socket.SetDeadline(time.Now().Add(duration))
	//if err != nil {
	//	panic(err)
	//}
	version := readByte(is)
	//if err != nil {
	//	panic(err)
	//}
	//duration, err = time.ParseDuration(fmt.Sprintf("%dms", t.messageFragmentTimeout))
	//err = socket.SetDeadline(time.Now().Add(duration))

	for eot != 0x80 {
		// read version
		if version != 3 {
			panic(errors.New("syntax error at beginning of RFC1006 header: version not equal to 3"))
		}
		// read reserved
		if readByte(is) != 0 {
			panic(errors.New("syntax errorat beginning of RFC1006 header: reserved not equal to 0"))
		}

		// read packet length
		packetLength = readShort(is) & 0xffff
		if packetLength <= 7 {
			panic(errors.New("syntax error: packet length parameter < 7"))
		}
		// read length indicator
		li = int(readByte(is) & 0xff)

		// read TPDU code
		b := readByte(is)
		tPduCode = int(b & 0xff)

		if tPduCode == 0xf0 {
			// Data Transfer (DT) Code

			if li != 2 {
				panic(errors.New("syntax error: LI field does not equal 2"))
			}

			// read EOT
			eot = int(readByte(is) & 0xff)
			if eot != 0 && eot != 0x80 {
				panic(errors.New("syntax error: eot wrong"))

			}

			//if (packetLength - 7 > tSduBuffer.limit() - tSduBuffer.position()) {
			//	panic(errors.New("tSduBuffer size is too small to hold the complete TSDU"))
			//}
			buf := make([]byte, packetLength-7)
			_, _ = is.Read(buf)
			tSduBuffer.Write(buf)

		} else if tPduCode == 0x80 {
			// Disconnect Request (DR)

			if li != 6 {
				panic(errors.New("syntax error: LI field does not equal 6"))
			}

			// check if the DST-REF field is set to the reference of the
			// receiving entity -> srcRef
			if readShort(is) != t.srcRef {
				panic(errors.New("syntax error: srcRef wrong"))

			}

			// check if the SRC-REF field is that of the entity sending
			// the DR
			if readShort(is) != t.dstRef {
				panic(errors.New("syntax error: dstRef wrong"))
			}

			// check the reason field, for class 0 only between 1 and 4
			reason := readByte(is) & 0xff
			if reason > 4 {
				panic(errors.New("syntax error: reason out of bound"))
			}
			panic(errors.New(fmt.Sprintf("Disconnect request. Reason:%b", reason)))

		} else if tPduCode == 0x70 {
			panic(errors.New("got TPDU error (ER) message"))

		} else {
			panic(errors.New("syntax error: unknown TPDU code"))
		}

		if eot != 0x80 {
			version = readByte(is)
		}

	}
}

/**
 * Starts a connection, sends a CR, waits for a CC and throws an IOException if not successful
 *
 * @throws IOException if an error occurs
 */
func (t *TConnection) startConnection() {
	os := t.os
	is := t.is
	// writeByte RFC 1006 Header
	writeByteByte(os, 0x03)
	writeByteByte(os, 0x00)
	// writeByte complete packet length
	variableLength := 3

	if t.TSelLocal != nil {
		variableLength += 2 + len(t.TSelLocal)
	}
	if t.TSelRemote != nil {
		variableLength += 2 + len(t.TSelRemote)
	}

	writeByteByteShort(os, 4+7+variableLength)
	// writing RFC 1006 Header finished

	// writeByte connection request (CR) TPDU (§13.3)

	// writeByte length indicator
	writeByteByte(os, byte(6+variableLength))

	// writeByte fixed part
	// writeByte CR CDT
	writeByteByte(os, 0xe0)
	// writeByte DST-REF
	writeByteByte(os, 0)
	writeByteByte(os, 0)
	// writeByte SRC-REF
	writeByteByteShort(os, t.srcRef)
	// writeByte class
	writeByteByte(os, 0)

	// writeByte variable part
	// writeByte proposed maximum TPDU Size
	writeByteByte(os, 0xc0)
	writeByteByte(os, 1)
	writeByteByte(os, byte(t.maxTPduSizeParam))

	if t.TSelRemote != nil {
		writeByteByte(os, 0xc2)
		writeByteByte(os, byte(len(t.TSelRemote)))
		_, err := os.Write(t.TSelRemote)
		if err != nil {
			panic(err)
		}
	}
	if t.TSelLocal != nil {
		writeByteByte(os, 0xc1)
		writeByteByte(os, byte(len(t.TSelLocal)))
		_, err := os.Write(t.TSelLocal)
		if err != nil {
			panic(err)
		}
	}

	err := os.Flush()
	if err != nil {
		panic(err)
	}

	//conn := *t.Socket

	//duration, err := time.ParseDuration(fmt.Sprintf("%dms", t.messageTimeout))
	//if err != nil {
	//	panic(err)
	//}
	//err = conn.SetDeadline(time.Now().Add(duration))
	//if err != nil {
	//	panic(err)
	//}

	var myByte byte
	var lengthIndicator int
	var parameterLength int
	if readByte(is) != 0x03 {
		panic(errors.New("io error"))
	}
	if readByte(is) != 0 {
		panic(errors.New("io error"))
	}
	// read packet length, but is not needed
	readShort(is)
	lengthIndicator = int(readByte(is) & 0xff)
	if (readByte(is) & 0xff) != 0xd0 {
		panic(errors.New("io error"))
	}
	// read the dstRef which is the srcRef for t end-point
	readShort(is)
	// read the srcRef which is the dstRef for t end-point
	t.dstRef = readShort(is) & 0xffff
	// read class
	if readByte(is) != 0 {
		panic(errors.New("io error"))
	}

	variableBytesRead := 0
	for lengthIndicator > (6 + variableBytesRead) {
		// read parameter code
		myByte = readByte(is)
		switch myByte & 0xff {
		case 0xc1:
			parameterLength = int(readByte(is) & 0xff)

			if t.TSelLocal == nil {
				t.TSelLocal = make([]byte, parameterLength)
				is.Read(t.TSelLocal)
			} else {
				for i := 0; i < parameterLength; i++ {
					read(is)
				}

			}
			variableBytesRead += 2 + parameterLength
			break
		case 0xc2:
			parameterLength = int(readByte(is) & 0xff)

			if t.TSelRemote == nil {
				t.TSelRemote = make([]byte, parameterLength)
				_, err = is.Read(t.TSelRemote)
				if err != nil {
					panic(err)
				}
			} else {
				for i := 0; i < parameterLength; i++ {
					read(is)
				}
			}
			variableBytesRead += 2 + parameterLength
			break

		case 0xc0:
			if readByte(is) != 1 {
				panic(errors.New("maxTPduSizeParam size is not equal to 1"))

			}
			myByte = readByte(is)
			if int(myByte&0xff) < 7 || int(myByte&0xff) > t.maxTPduSizeParam {
				panic(errors.New("maxTPduSizeParam out of bound"))

			} else {
				if int(myByte&0xff) < t.maxTPduSizeParam {
					t.maxTPduSizeParam = int(myByte & 0xff)
				}
			}
			variableBytesRead += 4
			break
		default:
			panic(errors.New("io error"))
		}
	}

}

/**
 * This function is called once a client has Connected to the server. It listens for a Connection
 * Request (CR). If this is successful it replies afterwards with a Connection Confirm (CC).
 * According to the norm a syntax error in the CR should be followed by an ER. This implementation
 * does not send an ER because it seems unnecessary.
 *
 * @throws IOException if an error occurs
 */
func (t *TConnection) listenForCR() {

	is := t.is
	os := t.os

	var lengthIndicator int
	var parameterLength int

	// start reading rfc 1006 header
	if read(is) != 0x03 {
		panic("io error")
	}
	if read(is) != 0 {
		panic("io error")
	}
	// read Packet length, but is not needed
	read(is)
	// reading rfc 1006 header finished

	lengthIndicator = int(read(is) & 0xff)
	// 0xe0 is the CR-Code

	if read(is) != 0xe0 {
		panic("io error")
	}
	// DST-REF needs to be 0 in a CR packet

	if readShort(is) != 0 {
		panic("io error")
	}

	// read the srcRef which is the dstRef for t end-point
	t.dstRef = readShort(is) & 0xffff

	// read class
	if read(is)&0xff != 0 {
		panic("io error")
	}

	variableBytesRead := 0
	for lengthIndicator > (6 + variableBytesRead) {
		// read parameter code
		switch readByte(is) & 0xff {
		case 0xc2:

			parameterLength = read(is) & 0xff

			if t.TSelLocal == nil {
				t.TSelLocal = make([]byte, parameterLength)
				_, _ = is.Read(t.TSelLocal)
			} else {
				if parameterLength != len(t.TSelLocal) {
					panic("local T-SElECTOR is wrong")
				}
				for i := 0; i < parameterLength; i++ {

					if (t.TSelLocal[i] & 0xff) != readByte(is) {
						panic("local T-SElECTOR is wrong")
					}
				}
			}
			variableBytesRead += 2 + parameterLength
			break
		case 0xc1:
			parameterLength = read(is) & 0xff

			if t.TSelRemote == nil {
				t.TSelRemote = make([]byte, parameterLength)
				_, _ = is.Read(t.TSelRemote)
			} else {
				if parameterLength != len(t.TSelRemote) {
					panic("remote T-SElECTOR is wrong")
				}
				for i := 0; i < parameterLength; i++ {
					if (int(t.TSelRemote[i] & 0xff)) != read(is) {
						panic("remote T-SElECTOR is wrong")
					}
				}

			}
			variableBytesRead += 2 + parameterLength
			break

		case 0xc0:
			if (readByte(is) & 0xff) != 1 {
				panic("io error")
			}
			newMaxTPDUSizeParam := int(readByte(is) & 0xff)
			if newMaxTPDUSizeParam < 7 || newMaxTPDUSizeParam > 16 {
				panic("maxTPDUSizeParam is out of bound")
			} else {
				if newMaxTPDUSizeParam < t.maxTPduSizeParam {
					t.maxTPduSizeParam = newMaxTPDUSizeParam
					t.maxTPduSize = getMaxTPDUSize(t.maxTPduSizeParam)
				}
			}
			variableBytesRead += 3
			break
		default:
			panic("io error")
		}
	}

	// writeByte RFC 1006 Header
	writeByteByte(os, 0x03)
	writeByteByte(os, 0x00)

	// writeByte complete packet length

	variableLength := 3

	if t.TSelLocal != nil {
		variableLength += 2 + len(t.TSelLocal)
	}
	if t.TSelRemote != nil {
		variableLength += 2 + len(t.TSelRemote)
	}

	writeByteByteShort(os, 4+7+variableLength)
	// writeByte connection request (CR) TPDU (§13.3)

	// writeByte length indicator
	writeByteByte(os, byte(6+variableLength))

	// writeByte fixed part
	// writeByte CC CDT
	writeByteByte(os, 0xd0)

	// writeByte DST-REF
	writeByteByteShort(os, t.dstRef)
	// writeByte SRC-REF
	writeByteByteShort(os, t.srcRef)
	// writeByte class
	writeByteByte(os, 0x00)

	// writeByte variable part
	if t.TSelLocal != nil {
		writeByteByte(os, 0xc2)
		writeByteByte(os, byte(len(t.TSelLocal)))
		_, err := os.Write(t.TSelLocal)
		if err != nil {
			panic(err)
		}
	}

	if t.TSelRemote != nil {
		writeByteByte(os, 0xc1)

		writeByteByte(os, byte(len(t.TSelRemote)))
		_, err := os.Write(t.TSelLocal)
		if err != nil {
			panic(err)
		}

	}
	// writeByte proposed maximum TPDU Size
	writeByteByte(os, 0xc0)

	writeByteByte(os, 1)

	writeByteByte(os, byte(t.maxTPduSizeParam))
	err := os.Flush()
	if err != nil {
		panic(err)
	}

}

/**
 * Calculates and returns the maximum TPDUSize. This is equal to 2^(maxTPDUSizeParam)
 *
 * @param maxTPDUSizeParam the size parameter
 * @return the maximum TPDU size
 */
func getMaxTPDUSize(maxTPDUSizeParam int) int {
	if maxTPDUSizeParam < 7 || maxTPDUSizeParam > 16 {
		panic("maxTPDUSizeParam is out of bound")
	}
	if maxTPDUSizeParam == 16 {
		return 65531
	} else {
		return int(math.Pow(2, float64(maxTPDUSizeParam)))
	}
}

func (t *TConnection) send(tsdus [][]byte, offsets []int, lengths []int) {

	//TODO
	os := t.os
	bytesLeft := 0
	for _, length := range lengths {
		bytesLeft += length
	}

	tsduOffset := 0
	byteArrayListIndex := 0
	var numBytesToWrite int
	lastPacket := false
	maxTSDUSize := t.maxTPduSize - 3
	for bytesLeft > 0 {
		if bytesLeft > maxTSDUSize {
			numBytesToWrite = maxTSDUSize
		} else {
			numBytesToWrite = bytesLeft
			lastPacket = true
		}

		// --writeByte RFC 1006 Header--
		// writeByte Version
		writeByteByte(os, 0x03)
		// writeByte reserved bits
		writeByteByte(os, 0)
		// writeByte packet Length
		writeByteByteShort(os, numBytesToWrite+7)

		// --writeByte 8073 Header--
		// writeByte Length Indicator of header
		writeByteByte(os, 0x02)
		// writeByte TPDU Code for DT Data
		writeByteByte(os, 0xf0)
		// writeByte TPDU-NR and EOT, TPDU-NR is always 0 for class 0
		if lastPacket {
			writeByteByte(os, 0x80)
		} else {
			writeByteByte(os, 0x00)
		}

		bytesLeft -= numBytesToWrite
		for numBytesToWrite > 0 {
			tsdu := tsdus[byteArrayListIndex]
			length := lengths[byteArrayListIndex]
			offset := offsets[byteArrayListIndex]

			tsduWriteLength := length - tsduOffset

			if numBytesToWrite > tsduWriteLength {
				_, err := os.Write(tsdu[offset+tsduOffset : offset+tsduOffset+tsduWriteLength])
				if err != nil {
					panic(err)
				}
				numBytesToWrite -= tsduWriteLength
				tsduOffset = 0
				byteArrayListIndex++
			} else {
				_, err := os.Write(tsdu[offset+tsduOffset : offset+tsduOffset+numBytesToWrite])
				if err != nil {
					panic(err)
				}
				if numBytesToWrite == tsduWriteLength {
					tsduOffset = 0
					byteArrayListIndex++
				} else {
					tsduOffset += numBytesToWrite
				}
				numBytesToWrite = 0
			}
		}

		err := os.Flush()
		if err != nil {
			panic(err)
		}
	}

}

func (t *TConnection) SendSingle(tsdu []byte, offset int, length int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	tsdus := make([][]byte, 0)
	tsdus = append(tsdus, tsdu)

	offsets :=
		make([]int, 0)
	offsets = append(offsets, offset)

	lengths := make([]int, 0)
	lengths = append(lengths, length)

	t.send(tsdus, offsets, lengths)

	return nil
}

/** This function sends a Disconnect Request but does not wait for a Disconnect Confirm. */
func (t *TConnection) disconnect() {
	defer func() {
		t.close()
	}()
	os := t.os
	// writeByte header for rfc
	// writeByte version
	writeByteByte(os, 0x03)
	// writeByte reserved
	writeByteByte(os, 0x00)
	// writeByte packet length
	writeByteByteShort(os, 4+7) // t does not include the variable part
	// which
	// contains additional user information for
	// disconnect

	// beginning of ISO 8073 header
	// writeByte LI
	writeByteByte(os, 0x06)

	// writeByte DR
	writeByteByte(os, 0x80)

	// writeByte DST-REF
	writeByteByteShort(os, t.dstRef)

	// writeByte SRC-REF
	writeByteByteShort(os, t.srcRef)

	// writeByte reason - 0x00 corresponds to reason not specified. Can
	// writeByte
	// the reasons as case structure, but need input from client
	writeByteByte(os, 0x00)

	err := os.Flush()
	if err != nil {
		panic(err)

	}

}

/** Will close the TCP connection if It's still open and free any resources of this connection. */
func (t *TConnection) close() {
	closed := t.Closed
	serverThread := t.serverThread
	if !closed {
		closed = true
		// will also close socket
		t.os = nil
		t.is = nil
		if serverThread != nil {
			serverThread.connectionClosedSignal()
		}
	}
}
