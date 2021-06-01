package goprisma
/*
#cgo darwin,!arm64 LDFLAGS: -L"${SRCDIR}/lib/darwin" -Wl,-rpath,"${SRCDIR}/lib/darwin" -lquery_engine_c_api
#cgo darwin,arm64 LDFLAGS: -L"${SRCDIR}/lib/darwin-aarch64" -Wl,-rpath,"${SRCDIR}/lib/darwin-aarch64" -lquery_engine_c_api
#cgo linux LDFLAGS: -L"${SRCDIR}/lib/linux" -Wl,-rpath,"${SRCDIR}/lib/linux" -lquery_engine_c_api
#cgo windows LDFLAGS: -L"${SRCDIR}/lib/windows" -Wl,-rpath,"${SRCDIR}/lib/windows" -lquery_engine_c_api
*/
import "C"