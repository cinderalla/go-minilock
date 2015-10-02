package minilock

import (
  "testing"
  //"fmt"
  //"github.com/cathalgarvey/go-minilock/minilockkeys"
  "github.com/cathalgarvey/go-minilock/minilockutils"
)

func Test_ParseMinilockFile(t *testing.T) {
  testcase, err := Asset("binary_samples/mye.go.minilock")
  if err != nil {
    t.Fatal("Couldn't load test binary asset.")
  }
  expected_plaintext, err := Asset("binary_samples/mye.go")
  if err != nil {
    t.Fatal("Couldn't load test binary asset.")
  }
  header, ciphertext, err := ParseFileContents(testcase)
  if err != nil {
    t.Fatal("Failed to parse testcase.")
  }
  //fmt.Println(header)
  // Either test1 or test2 should be able to decrypt but only one is the sender..
  test1 := testKey1
  senderID, filename, contents, err := header.DecryptContents(ciphertext, test1)
  if err != nil {
    t.Fatal("Failed to decrypt with testKey1: "+err.Error())
  }
  if senderID != "xjjCm44Nuj4DyTBuzguJ1d7K6EdP2TWRYzsqiiAbfcGTr" {
    t.Error("SenderID was expected to be 'xjjCm44Nuj4DyTBuzguJ1d7K6EdP2TWRYzsqiiAbfcGTr' but was: "+ senderID)
  }
  if filename != "mye.go" {
    t.Error("Filename returned should have been 'mye.go', was: "+filename)
  }
  if !minilockutils.CmpSlices(contents, expected_plaintext) {
    t.Error("Plaintext did not match expected plaintext.")
  }
  senderID2, filename2, contents2, err := DecryptFileContents(testcase, test1)
  if err != nil {
    t.Fatal("Failed to decrypt on second try with testKey1: "+err.Error())
  }
  if senderID != senderID2 {
    t.Error("Inconsistency between senderID returned by DecryptFileContents and manual parsing/header decryption.")
  }
  if filename != filename2 {
    t.Error("Inconsistency between filename returned by DecryptFileContents and manual parsing/header decryption.")
  }
  if !minilockutils.CmpSlices(contents2, expected_plaintext) {
    t.Error("Plaintext did not match expected plaintext.")
  }
}


// func (self *miniLockv1Header) ExtractDecryptInfo(recipientKey *minilockkeys.NaClKeypair) (nonce []byte, DI *DecryptInfoEntry, err error) {
// func (self *miniLockv1Header) ExtractFileInfo(recipientKey *minilockkeys.NaClKeypair) (*FileInfo, error) {
// func (self *miniLockv1Header) DecryptContents(ciphertext []byte, recipientKey *minilockkeys.NaClKeypair) (senderID, filename string, contents []byte, err error) {
