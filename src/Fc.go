package src

const (
	// The following FCs are not part of this enum because they are not really
	// FCs and only defined in part 8-1:
	// RP (report), LG (log), BR (buffered report), GO, GS, MS, US

	// FCs according to IEC 61850-7-2:
	/** Status information */
	ST = "ST"
	/** Measurands - analogue values */
	MX = "MX"
	/** Setpoint */
	SP = "SP"
	/** Substitution */
	SV = "SV"
	/** Configuration */
	CF = "CF"
	/** Description */
	DC = "DC"
	/** Setting group */
	SG = "SG"
	/** Setting group editable */
	SE = "SE"
	/** Service response / Service tracking */
	SR = "SR"
	/** Operate received */
	OR = "OR"
	/** Blocking */
	BL = "BL"
	/** Extended definition */
	EX = "EX"
	/** Control, deprecated but kept here for backward compatibility */
	CO = "CO"
	/** Unbuffered Reporting */
	RP = "RP"
	/** Buffered Reporting */
	BR = "BR"
)
