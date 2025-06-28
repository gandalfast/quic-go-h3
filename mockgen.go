//go:build gomock || generate

package quic

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_send_conn_test.go github.com/gandalfast/quic-go-h3 SendConn"
type SendConn = sendConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_raw_conn_test.go github.com/gandalfast/quic-go-h3 RawConn"
type RawConn = rawConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_sender_test.go github.com/gandalfast/quic-go-h3 Sender"
type Sender = sender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_stream_sender_test.go github.com/gandalfast/quic-go-h3 StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_stream_control_frame_getter_test.go github.com/gandalfast/quic-go-h3 StreamControlFrameGetter"
type StreamControlFrameGetter = streamControlFrameGetter

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_stream_frame_getter_test.go github.com/gandalfast/quic-go-h3 StreamFrameGetter"
type StreamFrameGetter = streamFrameGetter

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_frame_source_test.go github.com/gandalfast/quic-go-h3 FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_ack_frame_source_test.go github.com/gandalfast/quic-go-h3 AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_sealing_manager_test.go github.com/gandalfast/quic-go-h3 SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_unpacker_test.go github.com/gandalfast/quic-go-h3 Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_packer_test.go github.com/gandalfast/quic-go-h3 Packer"
type Packer = packer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_mtu_discoverer_test.go github.com/gandalfast/quic-go-h3 MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_conn_runner_test.go github.com/gandalfast/quic-go-h3 ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/gandalfast/quic-go-h3 -destination mock_packet_handler_test.go github.com/gandalfast/quic-go-h3 PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/gandalfast/quic-go-h3 -self_package github.com/gandalfast/quic-go-h3 -destination mock_packetconn_test.go net PacketConn"
