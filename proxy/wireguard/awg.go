package wireguard

import "strings"

func appendAwgIPC(cfg *strings.Builder, awg *AmneziaParameters) {
	if awg == nil {
		return
	}
	if awg.Jc != "" {
		cfg.WriteString("jc=" + awg.Jc + "\n")
	}
	if awg.Jmin != "" {
		cfg.WriteString("jmin=" + awg.Jmin + "\n")
	}
	if awg.Jmax != "" {
		cfg.WriteString("jmax=" + awg.Jmax + "\n")
	}
	if awg.S1 != "" {
		cfg.WriteString("s1=" + awg.S1 + "\n")
	}
	if awg.S2 != "" {
		cfg.WriteString("s2=" + awg.S2 + "\n")
	}
	if awg.S3 != "" {
		cfg.WriteString("s3=" + awg.S3 + "\n")
	}
	if awg.S4 != "" {
		cfg.WriteString("s4=" + awg.S4 + "\n")
	}
	if awg.H1 != "" {
		cfg.WriteString("h1=" + awg.H1 + "\n")
	}
	if awg.H2 != "" {
		cfg.WriteString("h2=" + awg.H2 + "\n")
	}
	if awg.H3 != "" {
		cfg.WriteString("h3=" + awg.H3 + "\n")
	}
	if awg.H4 != "" {
		cfg.WriteString("h4=" + awg.H4 + "\n")
	}
	if awg.I1 != "" {
		cfg.WriteString("i1=" + awg.I1 + "\n")
	}
	if awg.I2 != "" {
		cfg.WriteString("i2=" + awg.I2 + "\n")
	}
	if awg.I3 != "" {
		cfg.WriteString("i3=" + awg.I3 + "\n")
	}
	if awg.I4 != "" {
		cfg.WriteString("i4=" + awg.I4 + "\n")
	}
	if awg.I5 != "" {
		cfg.WriteString("i5=" + awg.I5 + "\n")
	}
}
