package main

import (
	"github.com/SSripilaipong/muto/common/rslt"

	"github.com/SSripilaipong/muto-play/common"
)

type cliReaderWrapper struct {
	*common.BufferedReader
}

func newCliReaderWrapper(bufferedReader *common.BufferedReader) cliReaderWrapper {
	return cliReaderWrapper{BufferedReader: bufferedReader}
}

func (wrapper cliReaderWrapper) Read() rslt.Of[string] {
	return rslt.New(wrapper.ReadLine())
}
