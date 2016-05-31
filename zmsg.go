// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate from msg_generate.go

package dns

//import (
//"encoding/base64"
//"net"
//)

// pack*() functions

func (rr *A) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packDataA(rr.A, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *AAAA) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packDataAAAA(rr.AAAA, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *CNAME) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Target, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *DNAME) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Target, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *HINFO) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packString(rr.Cpu, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packString(rr.Os, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *L32) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Preference, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packDataA(rr.Locator32, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *LOC) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint8(rr.Version, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint8(rr.Size, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint8(rr.HorizPre, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint8(rr.VertPre, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(rr.Latitude, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(rr.Longitude, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(rr.Altitude, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MB) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Mb, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MD) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Md, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MF) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Mf, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MG) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Mg, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MR) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Mr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MX) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Preference, msg, off)
	if err != nil {
		return off, err
	}
	off, err = PackDomainName(rr.Mx, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *NID) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Preference, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint64(rr.NodeID, msg, off, false)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *NS) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Ns, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *PTR) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Ptr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *RP) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = PackDomainName(rr.Mbox, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	off, err = PackDomainName(rr.Txt, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *SRV) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := rr.Hdr.pack(msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Priority, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(rr.Weight, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(rr.Port, msg, off)
	if err != nil {
		return off, err
	}
	off, err = PackDomainName(rr.Target, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

// unpack*() functions

func unpackA(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(A)
	rr.Hdr = h

	rr.A, off, err = unpackDataA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackAAAA(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(AAAA)
	rr.Hdr = h

	rr.AAAA, off, err = unpackDataAAAA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackCNAME(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(CNAME)
	rr.Hdr = h

	rr.Target, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackDNAME(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(DNAME)
	rr.Hdr = h

	rr.Target, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackHINFO(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(HINFO)
	rr.Hdr = h

	rr.Cpu, off, err = unpackString(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Os, off, err = unpackString(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackL32(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(L32)
	rr.Hdr = h

	rr.Preference, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Locator32, off, err = unpackDataA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackLOC(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(LOC)
	rr.Hdr = h

	rr.Version, off, err = unpackUint8(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Size, off, err = unpackUint8(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.HorizPre, off, err = unpackUint8(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.VertPre, off, err = unpackUint8(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Latitude, off, err = unpackUint32(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Longitude, off, err = unpackUint32(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Altitude, off, err = unpackUint32(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMB(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MB)
	rr.Hdr = h

	rr.Mb, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMD(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MD)
	rr.Hdr = h

	rr.Md, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMF(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MF)
	rr.Hdr = h

	rr.Mf, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMG(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MG)
	rr.Hdr = h

	rr.Mg, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMR(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MR)
	rr.Hdr = h

	rr.Mr, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMX(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MX)
	rr.Hdr = h

	rr.Preference, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Mx, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackNID(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(NID)
	rr.Hdr = h

	rr.Preference, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.NodeID, off, err = unpackUint64(msg, off, false)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackNS(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(NS)
	rr.Hdr = h

	rr.Ns, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackPTR(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(PTR)
	rr.Hdr = h

	rr.Ptr, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackRP(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(RP)
	rr.Hdr = h

	rr.Mbox, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Txt, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackSRV(h RR_Header, msg []byte, off int) (RR, int, error) {
	if noRdata(h) {
		return nil, off, nil
	}
	var err error
	rr := new(SRV)
	rr.Hdr = h

	rr.Priority, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Weight, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Port, off, err = unpackUint16(msg, off)
	if err != nil {
		return rr, off, err
	}
	if off == len(msg) {
		return rr, off, nil
	}
	rr.Target, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}
