#!/bin/bash -eu

export CXX="${CXX} -lresolv" # required by Go 1.20

compile_go_fuzzer github.com/gandalfast/quic-go-h3/fuzzing/frames Fuzz frame_fuzzer
compile_go_fuzzer github.com/gandalfast/quic-go-h3/fuzzing/header Fuzz header_fuzzer
compile_go_fuzzer github.com/gandalfast/quic-go-h3/fuzzing/transportparameters Fuzz transportparameter_fuzzer
compile_go_fuzzer github.com/gandalfast/quic-go-h3/fuzzing/tokens Fuzz token_fuzzer
compile_go_fuzzer github.com/gandalfast/quic-go-h3/fuzzing/handshake Fuzz handshake_fuzzer
