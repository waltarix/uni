//go:generate go run gen.go

// Package unidata contains information about Unicode characters.
package unidata

// Codepoint is a single codepoint.
type Codepoint struct {
	Width, Cat uint8
	Codepoint  uint32
	Name       string
}

// Emoji is an emoji sequence.
type Emoji struct {
	Codepoints            []uint32
	Name, Group, Subgroup string
	SkinTones             bool
	Genders               int
}

const (
	GenderNone = 0
	GenderSign = 1
	GenderRole = 2
)

// TODO: Some ZJW sequences contain too many ZJW joiners, e.g.:
// 👩‍❤‍️‍💋‍👨
// 👩‍❤️‍💋‍👨 E2.0 kiss: woman, man
func (e Emoji) String() string {
	var c string

	// Flags
	// 1F1FF 1F1FC                                 # 🇿🇼 E2.0 flag: Zimbabwe
	// 1F3F4 E0067 E0062 E0065 E006E E0067 E007F   # 🏴󠁧󠁢󠁥󠁮󠁧󠁿 E5.0 flag: England
	if (e.Codepoints[0] >= 0x1f1e6 && e.Codepoints[0] <= 0x1f1ff) ||
		(len(e.Codepoints) > 1 && e.Codepoints[1] == 0xe0067) {
		for _, cp := range e.Codepoints {
			c += string(cp)
		}
		return c
	}

	// Keycap: join with 0xfe0f

	for i, cp := range e.Codepoints {
		c += string(cp)

		// Don't add ZWJ as last item.
		if i == len(e.Codepoints)-1 {
			continue
		}
		switch e.Codepoints[i+1] {

		// Never add ZWJ before variation selector or skin tone.
		case 0xfe0f, 0x1f3fb, 0x1f3fc, 0x1f3fd, 0x1f3fe, 0x1f3ff:
			continue

		// Keycaps
		case 0x20e3:
			continue
		}

		c += "\u200d"
	}
	return c
}

const (
	WidthAmbiguous = uint8(iota) // Ambiguous, A
	WidthFullWidth               // FullWidth, F
	WidthHalfWidth               // Halfwidth, H
	WidthNarrow                  // Narrow, N
	WidthNeutral                 // Neutral (Not East Asian), Na
	WidthWide                    // Wide, W
)

// http://www.unicode.org/reports/tr44/#General_Category_Values
const (
	CatUppercaseLetter      = uint8(iota) // Lu – an uppercase letter
	CatLowercaseLetter                    // Ll – a lowercase letter
	CatTitlecaseLetter                    // Lt – a digraphic character, with first part uppercase
	CatCasedLetter                        // LC – Lu | Ll | Lt
	CatModifierLetter                     // Lm – a modifier letter
	CatOtherLetter                        // Lo – other letters, including syllables and ideographs
	CatLetter                             // L  – Lu | Ll | Lt | Lm | Lo
	CatNonspacingMark                     // Mn – a nonspacing combining mark (zero advance width)
	CatSpacingMark                        // Mc – a spacing combining mark (positive advance width)
	CatEnclosingMark                      // Me – an enclosing combining mark
	CatMark                               // M  – Mn | Mc | Me
	CatDecimalNumber                      // Nd – a decimal digit
	CatLetterNumber                       // Nl – a letterlike numeric character
	CatOtherNumber                        // No – a numeric character of other type
	CatNumber                             // N  – Nd | Nl | No
	CatConnectorPunctuation               // Pc – a connecting punctuation mark, like a tie
	CatDashPunctuation                    // Pd – a dash or hyphen punctuation mark
	CatOpenPunctuation                    // Ps – an opening punctuation mark (of a pair)
	CatClosePunctuation                   // Pe – a closing punctuation mark (of a pair)
	CatInitialPunctuation                 // Pi – an initial quotation mark
	CatFinalPunctuation                   // Pf – a final quotation mark
	CatOtherPunctuation                   // Po – a punctuation mark of other type
	CatPunctuation                        // P  – Pc | Pd | Ps | Pe | Pi | Pf | Po
	CatMathSymbol                         // Sm – a symbol of mathematical use
	CatCurrencySymbol                     // Sc – a currency sign
	CatModifierSymbol                     // Sk – a non-letterlike modifier symbol
	CatOtherSymbol                        // So – a symbol of other type
	CatSymbol                             // S  – Sm | Sc | Sk | So
	CatSpaceSeparator                     // Zs – a space character (of various non-zero widths)
	CatLineSeparator                      // Zl – U+2028 LINE SEPARATOR only
	CatParagraphSeparator                 // Zp – U+2029 PARAGRAPH SEPARATOR only
	CatSeparator                          // Z  – Zs | Zl | Zp
	CatControl                            // Cc – a C0 or C1 control code
	CatFormat                             // Cf – a format control character
	CatSurrogate                          // Cs – a surrogate code point
	CatPrivateUse                         // Co – a private-use character
	CatUnassigned                         // Cn – a reserved unassigned code point or a noncharacter
	CatOther                              // C  – Cc | Cf | Cs | Co | Cn
)

// https://www.unicode.org/Public/UCD/latest/ucd/Blocks.txt
// TODO: generate this from the data file.
var (
	Blocks = map[string][]uint32{
		"Basic Latin":                           {0x0000, 0x007F},
		"Latin-1 Supplement":                    {0x0080, 0x00FF},
		"Latin Extended-A":                      {0x0100, 0x017F},
		"Latin Extended-B":                      {0x0180, 0x024F},
		"IPA Extensions":                        {0x0250, 0x02AF},
		"Spacing Modifier Letters":              {0x02B0, 0x02FF},
		"Combining Diacritical Marks":           {0x0300, 0x036F},
		"Greek and Coptic":                      {0x0370, 0x03FF},
		"Cyrillic":                              {0x0400, 0x04FF},
		"Cyrillic Supplement":                   {0x0500, 0x052F},
		"Armenian":                              {0x0530, 0x058F},
		"Hebrew":                                {0x0590, 0x05FF},
		"Arabic":                                {0x0600, 0x06FF},
		"Syriac":                                {0x0700, 0x074F},
		"Arabic Supplement":                     {0x0750, 0x077F},
		"Thaana":                                {0x0780, 0x07BF},
		"NKo":                                   {0x07C0, 0x07FF},
		"Samaritan":                             {0x0800, 0x083F},
		"Mandaic":                               {0x0840, 0x085F},
		"Syriac Supplement":                     {0x0860, 0x086F},
		"Arabic Extended-A":                     {0x08A0, 0x08FF},
		"Devanagari":                            {0x0900, 0x097F},
		"Bengali":                               {0x0980, 0x09FF},
		"Gurmukhi":                              {0x0A00, 0x0A7F},
		"Gujarati":                              {0x0A80, 0x0AFF},
		"Oriya":                                 {0x0B00, 0x0B7F},
		"Tamil":                                 {0x0B80, 0x0BFF},
		"Telugu":                                {0x0C00, 0x0C7F},
		"Kannada":                               {0x0C80, 0x0CFF},
		"Malayalam":                             {0x0D00, 0x0D7F},
		"Sinhala":                               {0x0D80, 0x0DFF},
		"Thai":                                  {0x0E00, 0x0E7F},
		"Lao":                                   {0x0E80, 0x0EFF},
		"Tibetan":                               {0x0F00, 0x0FFF},
		"Myanmar":                               {0x1000, 0x109F},
		"Georgian":                              {0x10A0, 0x10FF},
		"Hangul Jamo":                           {0x1100, 0x11FF},
		"Ethiopic":                              {0x1200, 0x137F},
		"Ethiopic Supplement":                   {0x1380, 0x139F},
		"Cherokee":                              {0x13A0, 0x13FF},
		"Unified Canadian Aboriginal Syllabics": {0x1400, 0x167F},
		"Ogham":                                 {0x1680, 0x169F},
		"Runic":                                 {0x16A0, 0x16FF},
		"Tagalog":                               {0x1700, 0x171F},
		"Hanunoo":                               {0x1720, 0x173F},
		"Buhid":                                 {0x1740, 0x175F},
		"Tagbanwa":                              {0x1760, 0x177F},
		"Khmer":                                 {0x1780, 0x17FF},
		"Mongolian":                             {0x1800, 0x18AF},
		"Unified Canadian Aboriginal Syllabics Extended": {0x18B0, 0x18FF},
		"Limbu":                                  {0x1900, 0x194F},
		"Tai Le":                                 {0x1950, 0x197F},
		"New Tai Lue":                            {0x1980, 0x19DF},
		"Khmer Symbols":                          {0x19E0, 0x19FF},
		"Buginese":                               {0x1A00, 0x1A1F},
		"Tai Tham":                               {0x1A20, 0x1AAF},
		"Combining Diacritical Marks Extended":   {0x1AB0, 0x1AFF},
		"Balinese":                               {0x1B00, 0x1B7F},
		"Sundanese":                              {0x1B80, 0x1BBF},
		"Batak":                                  {0x1BC0, 0x1BFF},
		"Lepcha":                                 {0x1C00, 0x1C4F},
		"Ol Chiki":                               {0x1C50, 0x1C7F},
		"Cyrillic Extended-C":                    {0x1C80, 0x1C8F},
		"Georgian Extended":                      {0x1C90, 0x1CBF},
		"Sundanese Supplement":                   {0x1CC0, 0x1CCF},
		"Vedic Extensions":                       {0x1CD0, 0x1CFF},
		"Phonetic Extensions":                    {0x1D00, 0x1D7F},
		"Phonetic Extensions Supplement":         {0x1D80, 0x1DBF},
		"Combining Diacritical Marks Supplement": {0x1DC0, 0x1DFF},
		"Latin Extended Additional":              {0x1E00, 0x1EFF},
		"Greek Extended":                         {0x1F00, 0x1FFF},
		"General Punctuation":                    {0x2000, 0x206F},
		"Superscripts and Subscripts":            {0x2070, 0x209F},
		"Currency Symbols":                       {0x20A0, 0x20CF},
		"Combining Diacritical Marks for Symbols": {0x20D0, 0x20FF},
		"Letterlike Symbols":                      {0x2100, 0x214F},
		"Number Forms":                            {0x2150, 0x218F},
		"Arrows":                                  {0x2190, 0x21FF},
		"Mathematical Operators":                  {0x2200, 0x22FF},
		"Miscellaneous Technical":                 {0x2300, 0x23FF},
		"Control Pictures":                        {0x2400, 0x243F},
		"Optical Character Recognition":           {0x2440, 0x245F},
		"Enclosed Alphanumerics":                  {0x2460, 0x24FF},
		"Box Drawing":                             {0x2500, 0x257F},
		"Block Elements":                          {0x2580, 0x259F},
		"Geometric Shapes":                        {0x25A0, 0x25FF},
		"Miscellaneous Symbols":                   {0x2600, 0x26FF},
		"Dingbats":                                {0x2700, 0x27BF},
		"Miscellaneous Mathematical Symbols-A":    {0x27C0, 0x27EF},
		"Supplemental Arrows-A":                   {0x27F0, 0x27FF},
		"Braille Patterns":                        {0x2800, 0x28FF},
		"Supplemental Arrows-B":                   {0x2900, 0x297F},
		"Miscellaneous Mathematical Symbols-B":    {0x2980, 0x29FF},
		"Supplemental Mathematical Operators":     {0x2A00, 0x2AFF},
		"Miscellaneous Symbols and Arrows":        {0x2B00, 0x2BFF},
		"Glagolitic":                              {0x2C00, 0x2C5F},
		"Latin Extended-C":                        {0x2C60, 0x2C7F},
		"Coptic":                                  {0x2C80, 0x2CFF},
		"Georgian Supplement":                     {0x2D00, 0x2D2F},
		"Tifinagh":                                {0x2D30, 0x2D7F},
		"Ethiopic Extended":                       {0x2D80, 0x2DDF},
		"Cyrillic Extended-A":                     {0x2DE0, 0x2DFF},
		"Supplemental Punctuation":                {0x2E00, 0x2E7F},
		"CJK Radicals Supplement":                 {0x2E80, 0x2EFF},
		"Kangxi Radicals":                         {0x2F00, 0x2FDF},
		"Ideographic Description Characters":      {0x2FF0, 0x2FFF},
		"CJK Symbols and Punctuation":             {0x3000, 0x303F},
		"Hiragana":                                {0x3040, 0x309F},
		"Katakana":                                {0x30A0, 0x30FF},
		"Bopomofo":                                {0x3100, 0x312F},
		"Hangul Compatibility Jamo":               {0x3130, 0x318F},
		"Kanbun":                                  {0x3190, 0x319F},
		"Bopomofo Extended":                       {0x31A0, 0x31BF},
		"CJK Strokes":                             {0x31C0, 0x31EF},
		"Katakana Phonetic Extensions":            {0x31F0, 0x31FF},
		"Enclosed CJK Letters and Months":         {0x3200, 0x32FF},
		"CJK Compatibility":                       {0x3300, 0x33FF},
		"CJK Unified Ideographs Extension A":      {0x3400, 0x4DBF},
		"Yijing Hexagram Symbols":                 {0x4DC0, 0x4DFF},
		"CJK Unified Ideographs":                  {0x4E00, 0x9FFF},
		"Yi Syllables":                            {0xA000, 0xA48F},
		"Yi Radicals":                             {0xA490, 0xA4CF},
		"Lisu":                                    {0xA4D0, 0xA4FF},
		"Vai":                                     {0xA500, 0xA63F},
		"Cyrillic Extended-B":                     {0xA640, 0xA69F},
		"Bamum":                                   {0xA6A0, 0xA6FF},
		"Modifier Tone Letters":                   {0xA700, 0xA71F},
		"Latin Extended-D":                        {0xA720, 0xA7FF},
		"Syloti Nagri":                            {0xA800, 0xA82F},
		"Common Indic Number Forms":               {0xA830, 0xA83F},
		"Phags-pa":                                {0xA840, 0xA87F},
		"Saurashtra":                              {0xA880, 0xA8DF},
		"Devanagari Extended":                     {0xA8E0, 0xA8FF},
		"Kayah Li":                                {0xA900, 0xA92F},
		"Rejang":                                  {0xA930, 0xA95F},
		"Hangul Jamo Extended-A":                  {0xA960, 0xA97F},
		"Javanese":                                {0xA980, 0xA9DF},
		"Myanmar Extended-B":                      {0xA9E0, 0xA9FF},
		"Cham":                                    {0xAA00, 0xAA5F},
		"Myanmar Extended-A":                      {0xAA60, 0xAA7F},
		"Tai Viet":                                {0xAA80, 0xAADF},
		"Meetei Mayek Extensions":                 {0xAAE0, 0xAAFF},
		"Ethiopic Extended-A":                     {0xAB00, 0xAB2F},
		"Latin Extended-E":                        {0xAB30, 0xAB6F},
		"Cherokee Supplement":                     {0xAB70, 0xABBF},
		"Meetei Mayek":                            {0xABC0, 0xABFF},
		"Hangul Syllables":                        {0xAC00, 0xD7AF},
		"Hangul Jamo Extended-B":                  {0xD7B0, 0xD7FF},
		"High Surrogates":                         {0xD800, 0xDB7F},
		"High Private Use Surrogates":             {0xDB80, 0xDBFF},
		"Low Surrogates":                          {0xDC00, 0xDFFF},
		"Private Use Area":                        {0xE000, 0xF8FF},
		"CJK Compatibility Ideographs":            {0xF900, 0xFAFF},
		"Alphabetic Presentation Forms":           {0xFB00, 0xFB4F},
		"Arabic Presentation Forms-A":             {0xFB50, 0xFDFF},
		"Variation Selectors":                     {0xFE00, 0xFE0F},
		"Vertical Forms":                          {0xFE10, 0xFE1F},
		"Combining Half Marks":                    {0xFE20, 0xFE2F},
		"CJK Compatibility Forms":                 {0xFE30, 0xFE4F},
		"Small Form Variants":                     {0xFE50, 0xFE6F},
		"Arabic Presentation Forms-B":             {0xFE70, 0xFEFF},
		"Halfwidth and Fullwidth Forms":           {0xFF00, 0xFFEF},
		"Specials":                                {0xFFF0, 0xFFFF},
		"Linear B Syllabary":                      {0x10000, 0x1007F},
		"Linear B Ideograms":                      {0x10080, 0x100FF},
		"Aegean Numbers":                          {0x10100, 0x1013F},
		"Ancient Greek Numbers":                   {0x10140, 0x1018F},
		"Ancient Symbols":                         {0x10190, 0x101CF},
		"Phaistos Disc":                           {0x101D0, 0x101FF},
		"Lycian":                                  {0x10280, 0x1029F},
		"Carian":                                  {0x102A0, 0x102DF},
		"Coptic Epact Numbers":                    {0x102E0, 0x102FF},
		"Old Italic":                              {0x10300, 0x1032F},
		"Gothic":                                  {0x10330, 0x1034F},
		"Old Permic":                              {0x10350, 0x1037F},
		"Ugaritic":                                {0x10380, 0x1039F},
		"Old Persian":                             {0x103A0, 0x103DF},
		"Deseret":                                 {0x10400, 0x1044F},
		"Shavian":                                 {0x10450, 0x1047F},
		"Osmanya":                                 {0x10480, 0x104AF},
		"Osage":                                   {0x104B0, 0x104FF},
		"Elbasan":                                 {0x10500, 0x1052F},
		"Caucasian Albanian":                      {0x10530, 0x1056F},
		"Linear A":                                {0x10600, 0x1077F},
		"Cypriot Syllabary":                       {0x10800, 0x1083F},
		"Imperial Aramaic":                        {0x10840, 0x1085F},
		"Palmyrene":                               {0x10860, 0x1087F},
		"Nabataean":                               {0x10880, 0x108AF},
		"Hatran":                                  {0x108E0, 0x108FF},
		"Phoenician":                              {0x10900, 0x1091F},
		"Lydian":                                  {0x10920, 0x1093F},
		"Meroitic Hieroglyphs":                    {0x10980, 0x1099F},
		"Meroitic Cursive":                        {0x109A0, 0x109FF},
		"Kharoshthi":                              {0x10A00, 0x10A5F},
		"Old South Arabian":                       {0x10A60, 0x10A7F},
		"Old North Arabian":                       {0x10A80, 0x10A9F},
		"Manichaean":                              {0x10AC0, 0x10AFF},
		"Avestan":                                 {0x10B00, 0x10B3F},
		"Inscriptional Parthian":                  {0x10B40, 0x10B5F},
		"Inscriptional Pahlavi":                   {0x10B60, 0x10B7F},
		"Psalter Pahlavi":                         {0x10B80, 0x10BAF},
		"Old Turkic":                              {0x10C00, 0x10C4F},
		"Old Hungarian":                           {0x10C80, 0x10CFF},
		"Hanifi Rohingya":                         {0x10D00, 0x10D3F},
		"Rumi Numeral Symbols":                    {0x10E60, 0x10E7F},
		"Old Sogdian":                             {0x10F00, 0x10F2F},
		"Sogdian":                                 {0x10F30, 0x10F6F},
		"Elymaic":                                 {0x10FE0, 0x10FFF},
		"Brahmi":                                  {0x11000, 0x1107F},
		"Kaithi":                                  {0x11080, 0x110CF},
		"Sora Sompeng":                            {0x110D0, 0x110FF},
		"Chakma":                                  {0x11100, 0x1114F},
		"Mahajani":                                {0x11150, 0x1117F},
		"Sharada":                                 {0x11180, 0x111DF},
		"Sinhala Archaic Numbers":                 {0x111E0, 0x111FF},
		"Khojki":                                  {0x11200, 0x1124F},
		"Multani":                                 {0x11280, 0x112AF},
		"Khudawadi":                               {0x112B0, 0x112FF},
		"Grantha":                                 {0x11300, 0x1137F},
		"Newa":                                    {0x11400, 0x1147F},
		"Tirhuta":                                 {0x11480, 0x114DF},
		"Siddham":                                 {0x11580, 0x115FF},
		"Modi":                                    {0x11600, 0x1165F},
		"Mongolian Supplement":                    {0x11660, 0x1167F},
		"Takri":                                   {0x11680, 0x116CF},
		"Ahom":                                    {0x11700, 0x1173F},
		"Dogra":                                   {0x11800, 0x1184F},
		"Warang Citi":                             {0x118A0, 0x118FF},
		"Nandinagari":                             {0x119A0, 0x119FF},
		"Zanabazar Square":                        {0x11A00, 0x11A4F},
		"Soyombo":                                 {0x11A50, 0x11AAF},
		"Pau Cin Hau":                             {0x11AC0, 0x11AFF},
		"Bhaiksuki":                               {0x11C00, 0x11C6F},
		"Marchen":                                 {0x11C70, 0x11CBF},
		"Masaram Gondi":                           {0x11D00, 0x11D5F},
		"Gunjala Gondi":                           {0x11D60, 0x11DAF},
		"Makasar":                                 {0x11EE0, 0x11EFF},
		"Tamil Supplement":                        {0x11FC0, 0x11FFF},
		"Cuneiform":                               {0x12000, 0x123FF},
		"Cuneiform Numbers and Punctuation":       {0x12400, 0x1247F},
		"Early Dynastic Cuneiform":                {0x12480, 0x1254F},
		"Egyptian Hieroglyphs":                    {0x13000, 0x1342F},
		"Egyptian Hieroglyph Format Controls":     {0x13430, 0x1343F},
		"Anatolian Hieroglyphs":                   {0x14400, 0x1467F},
		"Bamum Supplement":                        {0x16800, 0x16A3F},
		"Mro":                                     {0x16A40, 0x16A6F},
		"Bassa Vah":                               {0x16AD0, 0x16AFF},
		"Pahawh Hmong":                            {0x16B00, 0x16B8F},
		"Medefaidrin":                             {0x16E40, 0x16E9F},
		"Miao":                                    {0x16F00, 0x16F9F},
		"Ideographic Symbols and Punctuation":     {0x16FE0, 0x16FFF},
		"Tangut":                                  {0x17000, 0x187FF},
		"Tangut Components":                       {0x18800, 0x18AFF},
		"Kana Supplement":                         {0x1B000, 0x1B0FF},
		"Kana Extended-A":                         {0x1B100, 0x1B12F},
		"Small Kana Extension":                    {0x1B130, 0x1B16F},
		"Nushu":                                   {0x1B170, 0x1B2FF},
		"Duployan":                                {0x1BC00, 0x1BC9F},
		"Shorthand Format Controls":               {0x1BCA0, 0x1BCAF},
		"Byzantine Musical Symbols":               {0x1D000, 0x1D0FF},
		"Musical Symbols":                         {0x1D100, 0x1D1FF},
		"Ancient Greek Musical Notation":          {0x1D200, 0x1D24F},
		"Mayan Numerals":                          {0x1D2E0, 0x1D2FF},
		"Tai Xuan Jing Symbols":                   {0x1D300, 0x1D35F},
		"Counting Rod Numerals":                   {0x1D360, 0x1D37F},
		"Mathematical Alphanumeric Symbols":       {0x1D400, 0x1D7FF},
		"Sutton SignWriting":                      {0x1D800, 0x1DAAF},
		"Glagolitic Supplement":                   {0x1E000, 0x1E02F},
		"Nyiakeng Puachue Hmong":                  {0x1E100, 0x1E14F},
		"Wancho":                                  {0x1E2C0, 0x1E2FF},
		"Mende Kikakui":                           {0x1E800, 0x1E8DF},
		"Adlam":                                   {0x1E900, 0x1E95F},
		"Indic Siyaq Numbers":                     {0x1EC70, 0x1ECBF},
		"Ottoman Siyaq Numbers":                   {0x1ED00, 0x1ED4F},
		"Arabic Mathematical Alphabetic Symbols":  {0x1EE00, 0x1EEFF},
		"Mahjong Tiles":                           {0x1F000, 0x1F02F},
		"Domino Tiles":                            {0x1F030, 0x1F09F},
		"Playing Cards":                           {0x1F0A0, 0x1F0FF},
		"Enclosed Alphanumeric Supplement":        {0x1F100, 0x1F1FF},
		"Enclosed Ideographic Supplement":         {0x1F200, 0x1F2FF},
		"Miscellaneous Symbols and Pictographs":   {0x1F300, 0x1F5FF},
		"Emoticons":                               {0x1F600, 0x1F64F},
		"Ornamental Dingbats":                     {0x1F650, 0x1F67F},
		"Transport and Map Symbols":               {0x1F680, 0x1F6FF},
		"Alchemical Symbols":                      {0x1F700, 0x1F77F},
		"Geometric Shapes Extended":               {0x1F780, 0x1F7FF},
		"Supplemental Arrows-C":                   {0x1F800, 0x1F8FF},
		"Supplemental Symbols and Pictographs":    {0x1F900, 0x1F9FF},
		"Chess Symbols":                           {0x1FA00, 0x1FA6F},
		"Symbols and Pictographs Extended-A":      {0x1FA70, 0x1FAFF},
		"CJK Unified Ideographs Extension B":      {0x20000, 0x2A6DF},
		"CJK Unified Ideographs Extension C":      {0x2A700, 0x2B73F},
		"CJK Unified Ideographs Extension D":      {0x2B740, 0x2B81F},
		"CJK Unified Ideographs Extension E":      {0x2B820, 0x2CEAF},
		"CJK Unified Ideographs Extension F":      {0x2CEB0, 0x2EBEF},
		"CJK Compatibility Ideographs Supplement": {0x2F800, 0x2FA1F},
		"Tags":                             {0xE0000, 0xE007F},
		"Variation Selectors Supplement":   {0xE0100, 0xE01EF},
		"Supplementary Private Use Area-A": {0xF0000, 0xFFFFF},
		"Supplementary Private Use Area-B": {0x100000, 0x10FFFF},
	}

	Blockmap = make(map[string]string)
)

func init() {
	for k := range Blocks {
		Blockmap[CanonicalCategory(k)] = k
	}
}

var (
	Catmap = map[string]uint8{
		// Short-hand.
		"Lu": CatUppercaseLetter,
		"Ll": CatLowercaseLetter,
		"Lt": CatTitlecaseLetter,
		"LC": CatCasedLetter,
		"Lm": CatModifierLetter,
		"Lo": CatOtherLetter,
		"L":  CatLetter,
		"Mn": CatNonspacingMark,
		"Mc": CatSpacingMark,
		"Me": CatEnclosingMark,
		"M":  CatMark,
		"Nd": CatDecimalNumber,
		"Nl": CatLetterNumber,
		"No": CatOtherNumber,
		"N":  CatNumber,
		"Pc": CatConnectorPunctuation,
		"Pd": CatDashPunctuation,
		"Ps": CatOpenPunctuation,
		"Pe": CatClosePunctuation,
		"Pi": CatInitialPunctuation,
		"Pf": CatFinalPunctuation,
		"Po": CatOtherPunctuation,
		"P":  CatPunctuation,
		"Sm": CatMathSymbol,
		"Sc": CatCurrencySymbol,
		"Sk": CatModifierSymbol,
		"So": CatOtherSymbol,
		"S":  CatSymbol,
		"Zs": CatSpaceSeparator,
		"Zl": CatLineSeparator,
		"Zp": CatParagraphSeparator,
		"Z":  CatSeparator,
		"Cc": CatControl,
		"Cf": CatFormat,
		"Cs": CatSurrogate,
		"Co": CatPrivateUse,
		"Cn": CatUnassigned,
		"C":  CatOther,

		// Lower-case shorthand.
		"lu": CatUppercaseLetter,
		"ll": CatLowercaseLetter,
		"lt": CatTitlecaseLetter,
		"lc": CatCasedLetter,
		"lm": CatModifierLetter,
		"lo": CatOtherLetter,
		"l":  CatLetter,
		"mn": CatNonspacingMark,
		"mc": CatSpacingMark,
		"me": CatEnclosingMark,
		"m":  CatMark,
		"nd": CatDecimalNumber,
		"nl": CatLetterNumber,
		"no": CatOtherNumber,
		"n":  CatNumber,
		"pc": CatConnectorPunctuation,
		"pd": CatDashPunctuation,
		"ps": CatOpenPunctuation,
		"pe": CatClosePunctuation,
		"pi": CatInitialPunctuation,
		"pf": CatFinalPunctuation,
		"po": CatOtherPunctuation,
		"p":  CatPunctuation,
		"sm": CatMathSymbol,
		"sc": CatCurrencySymbol,
		"sk": CatModifierSymbol,
		"so": CatOtherSymbol,
		"s":  CatSymbol,
		"zs": CatSpaceSeparator,
		"zl": CatLineSeparator,
		"zp": CatParagraphSeparator,
		"z":  CatSeparator,
		"cc": CatControl,
		"cf": CatFormat,
		"cs": CatSurrogate,
		"co": CatPrivateUse,
		"cn": CatUnassigned,
		"c":  CatOther,

		// Full names, underscores.
		"uppercase_letter":      CatUppercaseLetter,
		"lowercase_letter":      CatLowercaseLetter,
		"titlecase_letter":      CatTitlecaseLetter,
		"cased_letter":          CatCasedLetter,
		"modifier_letter":       CatModifierLetter,
		"other_letter":          CatOtherLetter,
		"letter":                CatLetter,
		"nonspacing_mark":       CatNonspacingMark,
		"spacing_mark":          CatSpacingMark,
		"enclosing_mark":        CatEnclosingMark,
		"mark":                  CatMark,
		"decimal_number":        CatDecimalNumber,
		"letter_number":         CatLetterNumber,
		"other_number":          CatOtherNumber,
		"number":                CatNumber,
		"connector_punctuation": CatConnectorPunctuation,
		"dash_punctuation":      CatDashPunctuation,
		"open_punctuation":      CatOpenPunctuation,
		"close_punctuation":     CatClosePunctuation,
		"initial_punctuation":   CatInitialPunctuation,
		"final_punctuation":     CatFinalPunctuation,
		"other_punctuation":     CatOtherPunctuation,
		"punctuation":           CatPunctuation,
		"math_symbol":           CatMathSymbol,
		"currency_symbol":       CatCurrencySymbol,
		"modifier_symbol":       CatModifierSymbol,
		"other_symbol":          CatOtherSymbol,
		"symbol":                CatSymbol,
		"space_separator":       CatSpaceSeparator,
		"line_separator":        CatLineSeparator,
		"paragraph_separator":   CatParagraphSeparator,
		"separator":             CatSeparator,
		"control":               CatControl,
		"format":                CatFormat,
		"surrogate":             CatSurrogate,
		"private_use":           CatPrivateUse,
		"unassigned":            CatUnassigned,
		"other":                 CatOther,

		// Without underscore.
		"uppercaseletter":      CatUppercaseLetter,
		"lowercaseletter":      CatLowercaseLetter,
		"titlecaseletter":      CatTitlecaseLetter,
		"casedletter":          CatCasedLetter,
		"modifierletter":       CatModifierLetter,
		"otherletter":          CatOtherLetter,
		"nonspacingmark":       CatNonspacingMark,
		"spacingmark":          CatSpacingMark,
		"enclosingmark":        CatEnclosingMark,
		"decimalnumber":        CatDecimalNumber,
		"letternumber":         CatLetterNumber,
		"othernumber":          CatOtherNumber,
		"connectorpunctuation": CatConnectorPunctuation,
		"dashpunctuation":      CatDashPunctuation,
		"openpunctuation":      CatOpenPunctuation,
		"closepunctuation":     CatClosePunctuation,
		"initialpunctuation":   CatInitialPunctuation,
		"finalpunctuation":     CatFinalPunctuation,
		"otherpunctuation":     CatOtherPunctuation,
		"mathsymbol":           CatMathSymbol,
		"currencysymbol":       CatCurrencySymbol,
		"modifiersymbol":       CatModifierSymbol,
		"othersymbol":          CatOtherSymbol,
		"spaceseparator":       CatSpaceSeparator,
		"lineseparator":        CatLineSeparator,
		"paragraphseparator":   CatParagraphSeparator,
		"privateuse":           CatPrivateUse,
	}

	Catnames = map[uint8]string{
		CatUppercaseLetter:      "Uppercase_Letter",
		CatLowercaseLetter:      "Lowercase_Letter",
		CatTitlecaseLetter:      "Titlecase_Letter",
		CatCasedLetter:          "Cased_Letter",
		CatModifierLetter:       "Modifier_Letter",
		CatOtherLetter:          "Other_Letter",
		CatLetter:               "Letter",
		CatNonspacingMark:       "Nonspacing_Mark",
		CatSpacingMark:          "Spacing_Mark",
		CatEnclosingMark:        "Enclosing_Mark",
		CatMark:                 "Mark",
		CatDecimalNumber:        "Decimal_Number",
		CatLetterNumber:         "Letter_Number",
		CatOtherNumber:          "Other_Number",
		CatNumber:               "Number",
		CatConnectorPunctuation: "Connector_Punctuation",
		CatDashPunctuation:      "Dash_Punctuation",
		CatOpenPunctuation:      "Open_Punctuation",
		CatClosePunctuation:     "Close_Punctuation",
		CatInitialPunctuation:   "Initial_Punctuation",
		CatFinalPunctuation:     "Final_Punctuation",
		CatOtherPunctuation:     "Other_Punctuation",
		CatPunctuation:          "Punctuation",
		CatMathSymbol:           "Math_Symbol",
		CatCurrencySymbol:       "Currency_Symbol",
		CatModifierSymbol:       "Modifier_Symbol",
		CatOtherSymbol:          "Other_Symbol",
		CatSymbol:               "Symbol",
		CatSpaceSeparator:       "Space_Separator",
		CatLineSeparator:        "Line_Separator",
		CatParagraphSeparator:   "Paragraph_Separator",
		CatSeparator:            "Separator",
		CatControl:              "Control",
		CatFormat:               "Format",
		CatSurrogate:            "Surrogate",
		CatPrivateUse:           "Private_Use",
		CatUnassigned:           "Unassigned",
		CatOther:                "Other",
	}
)

var (
	ranges = [][]rune{
		{0x3400, 0x4DB5},
		{0x4E00, 0x9FEF},
		{0xAC00, 0xD7A3},
		{0xD800, 0xDB7F},
		{0xDB80, 0xDBFF},
		{0xDC00, 0xDFFF},
		{0xE000, 0xF8FF},
		{0x17000, 0x187F1},
		{0x20000, 0x2A6D6},
		{0x2A700, 0x2B734},
		{0x2B740, 0x2B81D},
		{0x2B820, 0x2CEA1},
		{0x2CEB0, 0x2EBE0},
		{0xF0000, 0xFFFFD},
		{0x100000, 0x10FFFD},
	}

	rangeNames = []string{
		"<CJK Ideograph Extension A>",
		"<CJK Ideograph>",
		"<Hangul Syllable>",
		"<Non Private Use High Surrogate>",
		"<Private Use High Surrogate>",
		"<Low Surrogate>",
		"<Private Use>",
		"<Tangut Ideograph>",
		"<CJK Ideograph Extension B>",
		"<CJK Ideograph Extension C>",
		"<CJK Ideograph Extension D>",
		"<CJK Ideograph Extension E>",
		"<CJK Ideograph Extension F>",
		"<Plane 15 Private Use>",
		"<Plane 16 Private Use>",
	}
)
