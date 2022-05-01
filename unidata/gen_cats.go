// Code generated by gen.go; DO NOT EDIT

package unidata

// Unicode general categories (long names).
const (
	CatUnknown              = Category(iota)
	CatOther                // C   (Cc | Cf | Cn | Co | Cs)
	CatControl              // Cc
	CatFormat               // Cf
	CatUnassigned           // Cn
	CatPrivateUse           // Co
	CatSurrogate            // Cs
	CatLetter               // L   (Ll | Lm | Lo | Lt | Lu)
	CatCasedLetter          // LC  (Ll | Lt | Lu)
	CatLowercaseLetter      // Ll
	CatModifierLetter       // Lm
	CatOtherLetter          // Lo
	CatTitlecaseLetter      // Lt
	CatUppercaseLetter      // Lu
	CatMark                 // M   (Mc | Me | Mn)
	CatSpacingMark          // Mc
	CatEnclosingMark        // Me
	CatNonspacingMark       // Mn
	CatNumber               // N   (Nd | Nl | No)
	CatDecimalNumber        // Nd
	CatLetterNumber         // Nl
	CatOtherNumber          // No
	CatPunctuation          // P   (Pc | Pd | Pe | Pf | Pi | Po | Ps)
	CatConnectorPunctuation // Pc
	CatDashPunctuation      // Pd
	CatClosePunctuation     // Pe
	CatFinalPunctuation     // Pf
	CatInitialPunctuation   // Pi
	CatOtherPunctuation     // Po
	CatOpenPunctuation      // Ps
	CatSymbol               // S   (Sc | Sk | Sm | So)
	CatCurrencySymbol       // Sc
	CatModifierSymbol       // Sk
	CatMathSymbol           // Sm
	CatOtherSymbol          // So
	CatSeparator            // Z   (Zl | Zp | Zs)
	CatLineSeparator        // Zl
	CatParagraphSeparator   // Zp
	CatSpaceSeparator       // Zs
)

// Unicode general categories (short names).
const (
	CatC  = CatOther
	CatCc = CatControl
	CatCf = CatFormat
	CatCn = CatUnassigned
	CatCo = CatPrivateUse
	CatCs = CatSurrogate
	CatL  = CatLetter
	CatLC = CatCasedLetter
	CatLl = CatLowercaseLetter
	CatLm = CatModifierLetter
	CatLo = CatOtherLetter
	CatLt = CatTitlecaseLetter
	CatLu = CatUppercaseLetter
	CatM  = CatMark
	CatMc = CatSpacingMark
	CatMe = CatEnclosingMark
	CatMn = CatNonspacingMark
	CatN  = CatNumber
	CatNd = CatDecimalNumber
	CatNl = CatLetterNumber
	CatNo = CatOtherNumber
	CatP  = CatPunctuation
	CatPc = CatConnectorPunctuation
	CatPd = CatDashPunctuation
	CatPe = CatClosePunctuation
	CatPf = CatFinalPunctuation
	CatPi = CatInitialPunctuation
	CatPo = CatOtherPunctuation
	CatPs = CatOpenPunctuation
	CatS  = CatSymbol
	CatSc = CatCurrencySymbol
	CatSk = CatModifierSymbol
	CatSm = CatMathSymbol
	CatSo = CatOtherSymbol
	CatZ  = CatSeparator
	CatZl = CatLineSeparator
	CatZp = CatParagraphSeparator
	CatZs = CatSpaceSeparator
)

// Categories is a list of all categories.
var Categories = map[Category]struct {
	ShortName, Name string
	Include         []Category
}{
	CatOther:                {"C", "Other", []Category{CatCc, CatCf, CatCn, CatCo, CatCs}},
	CatControl:              {"Cc", "Control", nil},
	CatFormat:               {"Cf", "Format", nil},
	CatUnassigned:           {"Cn", "Unassigned", nil},
	CatPrivateUse:           {"Co", "Private_Use", nil},
	CatSurrogate:            {"Cs", "Surrogate", nil},
	CatLetter:               {"L", "Letter", []Category{CatLl, CatLm, CatLo, CatLt, CatLu}},
	CatCasedLetter:          {"LC", "Cased_Letter", []Category{CatLl, CatLt, CatLu}},
	CatLowercaseLetter:      {"Ll", "Lowercase_Letter", nil},
	CatModifierLetter:       {"Lm", "Modifier_Letter", nil},
	CatOtherLetter:          {"Lo", "Other_Letter", nil},
	CatTitlecaseLetter:      {"Lt", "Titlecase_Letter", nil},
	CatUppercaseLetter:      {"Lu", "Uppercase_Letter", nil},
	CatMark:                 {"M", "Mark", []Category{CatMc, CatMe, CatMn}},
	CatSpacingMark:          {"Mc", "Spacing_Mark", nil},
	CatEnclosingMark:        {"Me", "Enclosing_Mark", nil},
	CatNonspacingMark:       {"Mn", "Nonspacing_Mark", nil},
	CatNumber:               {"N", "Number", []Category{CatNd, CatNl, CatNo}},
	CatDecimalNumber:        {"Nd", "Decimal_Number", nil},
	CatLetterNumber:         {"Nl", "Letter_Number", nil},
	CatOtherNumber:          {"No", "Other_Number", nil},
	CatPunctuation:          {"P", "Punctuation", []Category{CatPc, CatPd, CatPe, CatPf, CatPi, CatPo, CatPs}},
	CatConnectorPunctuation: {"Pc", "Connector_Punctuation", nil},
	CatDashPunctuation:      {"Pd", "Dash_Punctuation", nil},
	CatClosePunctuation:     {"Pe", "Close_Punctuation", nil},
	CatFinalPunctuation:     {"Pf", "Final_Punctuation", nil},
	CatInitialPunctuation:   {"Pi", "Initial_Punctuation", nil},
	CatOtherPunctuation:     {"Po", "Other_Punctuation", nil},
	CatOpenPunctuation:      {"Ps", "Open_Punctuation", nil},
	CatSymbol:               {"S", "Symbol", []Category{CatSc, CatSk, CatSm, CatSo}},
	CatCurrencySymbol:       {"Sc", "Currency_Symbol", nil},
	CatModifierSymbol:       {"Sk", "Modifier_Symbol", nil},
	CatMathSymbol:           {"Sm", "Math_Symbol", nil},
	CatOtherSymbol:          {"So", "Other_Symbol", nil},
	CatSeparator:            {"Z", "Separator", []Category{CatZl, CatZp, CatZs}},
	CatLineSeparator:        {"Zl", "Line_Separator", nil},
	CatParagraphSeparator:   {"Zp", "Paragraph_Separator", nil},
	CatSpaceSeparator:       {"Zs", "Space_Separator", nil},
}
