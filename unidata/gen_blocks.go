// Code generated by gen.zsh; DO NOT EDIT

package unidata

// Unicode blocks
const (
	BlockUnknown = Block(iota)
	BlockBasicLatin
	BlockLatin1Supplement
	BlockLatinExtendedA
	BlockLatinExtendedB
	BlockIPAExtensions
	BlockSpacingModifierLetters
	BlockCombiningDiacriticalMarks
	BlockGreekandCoptic
	BlockCyrillic
	BlockCyrillicSupplement
	BlockArmenian
	BlockHebrew
	BlockArabic
	BlockSyriac
	BlockArabicSupplement
	BlockThaana
	BlockNKo
	BlockSamaritan
	BlockMandaic
	BlockSyriacSupplement
	BlockArabicExtendedB
	BlockArabicExtendedA
	BlockDevanagari
	BlockBengali
	BlockGurmukhi
	BlockGujarati
	BlockOriya
	BlockTamil
	BlockTelugu
	BlockKannada
	BlockMalayalam
	BlockSinhala
	BlockThai
	BlockLao
	BlockTibetan
	BlockMyanmar
	BlockGeorgian
	BlockHangulJamo
	BlockEthiopic
	BlockEthiopicSupplement
	BlockCherokee
	BlockUnifiedCanadianAboriginalSyllabics
	BlockOgham
	BlockRunic
	BlockTagalog
	BlockHanunoo
	BlockBuhid
	BlockTagbanwa
	BlockKhmer
	BlockMongolian
	BlockUnifiedCanadianAboriginalSyllabicsExtended
	BlockLimbu
	BlockTaiLe
	BlockNewTaiLue
	BlockKhmerSymbols
	BlockBuginese
	BlockTaiTham
	BlockCombiningDiacriticalMarksExtended
	BlockBalinese
	BlockSundanese
	BlockBatak
	BlockLepcha
	BlockOlChiki
	BlockCyrillicExtendedC
	BlockGeorgianExtended
	BlockSundaneseSupplement
	BlockVedicExtensions
	BlockPhoneticExtensions
	BlockPhoneticExtensionsSupplement
	BlockCombiningDiacriticalMarksSupplement
	BlockLatinExtendedAdditional
	BlockGreekExtended
	BlockGeneralPunctuation
	BlockSuperscriptsandSubscripts
	BlockCurrencySymbols
	BlockCombiningDiacriticalMarksforSymbols
	BlockLetterlikeSymbols
	BlockNumberForms
	BlockArrows
	BlockMathematicalOperators
	BlockMiscellaneousTechnical
	BlockControlPictures
	BlockOpticalCharacterRecognition
	BlockEnclosedAlphanumerics
	BlockBoxDrawing
	BlockBlockElements
	BlockGeometricShapes
	BlockMiscellaneousSymbols
	BlockDingbats
	BlockMiscellaneousMathematicalSymbolsA
	BlockSupplementalArrowsA
	BlockBraillePatterns
	BlockSupplementalArrowsB
	BlockMiscellaneousMathematicalSymbolsB
	BlockSupplementalMathematicalOperators
	BlockMiscellaneousSymbolsandArrows
	BlockGlagolitic
	BlockLatinExtendedC
	BlockCoptic
	BlockGeorgianSupplement
	BlockTifinagh
	BlockEthiopicExtended
	BlockCyrillicExtendedA
	BlockSupplementalPunctuation
	BlockCJKRadicalsSupplement
	BlockKangxiRadicals
	BlockIdeographicDescriptionCharacters
	BlockCJKSymbolsandPunctuation
	BlockHiragana
	BlockKatakana
	BlockBopomofo
	BlockHangulCompatibilityJamo
	BlockKanbun
	BlockBopomofoExtended
	BlockCJKStrokes
	BlockKatakanaPhoneticExtensions
	BlockEnclosedCJKLettersandMonths
	BlockCJKCompatibility
	BlockCJKUnifiedIdeographsExtensionA
	BlockYijingHexagramSymbols
	BlockCJKUnifiedIdeographs
	BlockYiSyllables
	BlockYiRadicals
	BlockLisu
	BlockVai
	BlockCyrillicExtendedB
	BlockBamum
	BlockModifierToneLetters
	BlockLatinExtendedD
	BlockSylotiNagri
	BlockCommonIndicNumberForms
	BlockPhagspa
	BlockSaurashtra
	BlockDevanagariExtended
	BlockKayahLi
	BlockRejang
	BlockHangulJamoExtendedA
	BlockJavanese
	BlockMyanmarExtendedB
	BlockCham
	BlockMyanmarExtendedA
	BlockTaiViet
	BlockMeeteiMayekExtensions
	BlockEthiopicExtendedA
	BlockLatinExtendedE
	BlockCherokeeSupplement
	BlockMeeteiMayek
	BlockHangulSyllables
	BlockHangulJamoExtendedB
	BlockHighSurrogates
	BlockHighPrivateUseSurrogates
	BlockLowSurrogates
	BlockPrivateUseArea
	BlockCJKCompatibilityIdeographs
	BlockAlphabeticPresentationForms
	BlockArabicPresentationFormsA
	BlockVariationSelectors
	BlockVerticalForms
	BlockCombiningHalfMarks
	BlockCJKCompatibilityForms
	BlockSmallFormVariants
	BlockArabicPresentationFormsB
	BlockHalfwidthandFullwidthForms
	BlockSpecials
	BlockLinearBSyllabary
	BlockLinearBIdeograms
	BlockAegeanNumbers
	BlockAncientGreekNumbers
	BlockAncientSymbols
	BlockPhaistosDisc
	BlockLycian
	BlockCarian
	BlockCopticEpactNumbers
	BlockOldItalic
	BlockGothic
	BlockOldPermic
	BlockUgaritic
	BlockOldPersian
	BlockDeseret
	BlockShavian
	BlockOsmanya
	BlockOsage
	BlockElbasan
	BlockCaucasianAlbanian
	BlockVithkuqi
	BlockLinearA
	BlockLatinExtendedF
	BlockCypriotSyllabary
	BlockImperialAramaic
	BlockPalmyrene
	BlockNabataean
	BlockHatran
	BlockPhoenician
	BlockLydian
	BlockMeroiticHieroglyphs
	BlockMeroiticCursive
	BlockKharoshthi
	BlockOldSouthArabian
	BlockOldNorthArabian
	BlockManichaean
	BlockAvestan
	BlockInscriptionalParthian
	BlockInscriptionalPahlavi
	BlockPsalterPahlavi
	BlockOldTurkic
	BlockOldHungarian
	BlockHanifiRohingya
	BlockRumiNumeralSymbols
	BlockYezidi
	BlockOldSogdian
	BlockSogdian
	BlockOldUyghur
	BlockChorasmian
	BlockElymaic
	BlockBrahmi
	BlockKaithi
	BlockSoraSompeng
	BlockChakma
	BlockMahajani
	BlockSharada
	BlockSinhalaArchaicNumbers
	BlockKhojki
	BlockMultani
	BlockKhudawadi
	BlockGrantha
	BlockNewa
	BlockTirhuta
	BlockSiddham
	BlockModi
	BlockMongolianSupplement
	BlockTakri
	BlockAhom
	BlockDogra
	BlockWarangCiti
	BlockDivesAkuru
	BlockNandinagari
	BlockZanabazarSquare
	BlockSoyombo
	BlockUnifiedCanadianAboriginalSyllabicsExtendedA
	BlockPauCinHau
	BlockBhaiksuki
	BlockMarchen
	BlockMasaramGondi
	BlockGunjalaGondi
	BlockMakasar
	BlockLisuSupplement
	BlockTamilSupplement
	BlockCuneiform
	BlockCuneiformNumbersandPunctuation
	BlockEarlyDynasticCuneiform
	BlockCyproMinoan
	BlockEgyptianHieroglyphs
	BlockEgyptianHieroglyphFormatControls
	BlockAnatolianHieroglyphs
	BlockBamumSupplement
	BlockMro
	BlockTangsa
	BlockBassaVah
	BlockPahawhHmong
	BlockMedefaidrin
	BlockMiao
	BlockIdeographicSymbolsandPunctuation
	BlockTangut
	BlockTangutComponents
	BlockKhitanSmallScript
	BlockTangutSupplement
	BlockKanaExtendedB
	BlockKanaSupplement
	BlockKanaExtendedA
	BlockSmallKanaExtension
	BlockNushu
	BlockDuployan
	BlockShorthandFormatControls
	BlockZnamennyMusicalNotation
	BlockByzantineMusicalSymbols
	BlockMusicalSymbols
	BlockAncientGreekMusicalNotation
	BlockMayanNumerals
	BlockTaiXuanJingSymbols
	BlockCountingRodNumerals
	BlockMathematicalAlphanumericSymbols
	BlockSuttonSignWriting
	BlockLatinExtendedG
	BlockGlagoliticSupplement
	BlockNyiakengPuachueHmong
	BlockToto
	BlockWancho
	BlockEthiopicExtendedB
	BlockMendeKikakui
	BlockAdlam
	BlockIndicSiyaqNumbers
	BlockOttomanSiyaqNumbers
	BlockArabicMathematicalAlphabeticSymbols
	BlockMahjongTiles
	BlockDominoTiles
	BlockPlayingCards
	BlockEnclosedAlphanumericSupplement
	BlockEnclosedIdeographicSupplement
	BlockMiscellaneousSymbolsandPictographs
	BlockEmoticons
	BlockOrnamentalDingbats
	BlockTransportandMapSymbols
	BlockAlchemicalSymbols
	BlockGeometricShapesExtended
	BlockSupplementalArrowsC
	BlockSupplementalSymbolsandPictographs
	BlockChessSymbols
	BlockSymbolsandPictographsExtendedA
	BlockSymbolsforLegacyComputing
	BlockCJKUnifiedIdeographsExtensionB
	BlockCJKUnifiedIdeographsExtensionC
	BlockCJKUnifiedIdeographsExtensionD
	BlockCJKUnifiedIdeographsExtensionE
	BlockCJKUnifiedIdeographsExtensionF
	BlockCJKCompatibilityIdeographsSupplement
	BlockCJKUnifiedIdeographsExtensionG
	BlockTags
	BlockVariationSelectorsSupplement
	BlockSupplementaryPrivateUseAreaA
	BlockSupplementaryPrivateUseAreaB
)

// Blocks is a list of all Unicode blocks.
var Blocks = map[Block]struct {
	Range [2]rune
	Name  string
}{
	BlockBasicLatin:                         {[2]rune{0x000000, 0x00007F}, "Basic Latin"},
	BlockLatin1Supplement:                   {[2]rune{0x000080, 0x0000FF}, "Latin-1 Supplement"},
	BlockLatinExtendedA:                     {[2]rune{0x000100, 0x00017F}, "Latin Extended-A"},
	BlockLatinExtendedB:                     {[2]rune{0x000180, 0x00024F}, "Latin Extended-B"},
	BlockIPAExtensions:                      {[2]rune{0x000250, 0x0002AF}, "IPA Extensions"},
	BlockSpacingModifierLetters:             {[2]rune{0x0002B0, 0x0002FF}, "Spacing Modifier Letters"},
	BlockCombiningDiacriticalMarks:          {[2]rune{0x000300, 0x00036F}, "Combining Diacritical Marks"},
	BlockGreekandCoptic:                     {[2]rune{0x000370, 0x0003FF}, "Greek and Coptic"},
	BlockCyrillic:                           {[2]rune{0x000400, 0x0004FF}, "Cyrillic"},
	BlockCyrillicSupplement:                 {[2]rune{0x000500, 0x00052F}, "Cyrillic Supplement"},
	BlockArmenian:                           {[2]rune{0x000530, 0x00058F}, "Armenian"},
	BlockHebrew:                             {[2]rune{0x000590, 0x0005FF}, "Hebrew"},
	BlockArabic:                             {[2]rune{0x000600, 0x0006FF}, "Arabic"},
	BlockSyriac:                             {[2]rune{0x000700, 0x00074F}, "Syriac"},
	BlockArabicSupplement:                   {[2]rune{0x000750, 0x00077F}, "Arabic Supplement"},
	BlockThaana:                             {[2]rune{0x000780, 0x0007BF}, "Thaana"},
	BlockNKo:                                {[2]rune{0x0007C0, 0x0007FF}, "NKo"},
	BlockSamaritan:                          {[2]rune{0x000800, 0x00083F}, "Samaritan"},
	BlockMandaic:                            {[2]rune{0x000840, 0x00085F}, "Mandaic"},
	BlockSyriacSupplement:                   {[2]rune{0x000860, 0x00086F}, "Syriac Supplement"},
	BlockArabicExtendedB:                    {[2]rune{0x000870, 0x00089F}, "Arabic Extended-B"},
	BlockArabicExtendedA:                    {[2]rune{0x0008A0, 0x0008FF}, "Arabic Extended-A"},
	BlockDevanagari:                         {[2]rune{0x000900, 0x00097F}, "Devanagari"},
	BlockBengali:                            {[2]rune{0x000980, 0x0009FF}, "Bengali"},
	BlockGurmukhi:                           {[2]rune{0x000A00, 0x000A7F}, "Gurmukhi"},
	BlockGujarati:                           {[2]rune{0x000A80, 0x000AFF}, "Gujarati"},
	BlockOriya:                              {[2]rune{0x000B00, 0x000B7F}, "Oriya"},
	BlockTamil:                              {[2]rune{0x000B80, 0x000BFF}, "Tamil"},
	BlockTelugu:                             {[2]rune{0x000C00, 0x000C7F}, "Telugu"},
	BlockKannada:                            {[2]rune{0x000C80, 0x000CFF}, "Kannada"},
	BlockMalayalam:                          {[2]rune{0x000D00, 0x000D7F}, "Malayalam"},
	BlockSinhala:                            {[2]rune{0x000D80, 0x000DFF}, "Sinhala"},
	BlockThai:                               {[2]rune{0x000E00, 0x000E7F}, "Thai"},
	BlockLao:                                {[2]rune{0x000E80, 0x000EFF}, "Lao"},
	BlockTibetan:                            {[2]rune{0x000F00, 0x000FFF}, "Tibetan"},
	BlockMyanmar:                            {[2]rune{0x001000, 0x00109F}, "Myanmar"},
	BlockGeorgian:                           {[2]rune{0x0010A0, 0x0010FF}, "Georgian"},
	BlockHangulJamo:                         {[2]rune{0x001100, 0x0011FF}, "Hangul Jamo"},
	BlockEthiopic:                           {[2]rune{0x001200, 0x00137F}, "Ethiopic"},
	BlockEthiopicSupplement:                 {[2]rune{0x001380, 0x00139F}, "Ethiopic Supplement"},
	BlockCherokee:                           {[2]rune{0x0013A0, 0x0013FF}, "Cherokee"},
	BlockUnifiedCanadianAboriginalSyllabics: {[2]rune{0x001400, 0x00167F}, "Unified Canadian Aboriginal Syllabics"},
	BlockOgham:                              {[2]rune{0x001680, 0x00169F}, "Ogham"},
	BlockRunic:                              {[2]rune{0x0016A0, 0x0016FF}, "Runic"},
	BlockTagalog:                            {[2]rune{0x001700, 0x00171F}, "Tagalog"},
	BlockHanunoo:                            {[2]rune{0x001720, 0x00173F}, "Hanunoo"},
	BlockBuhid:                              {[2]rune{0x001740, 0x00175F}, "Buhid"},
	BlockTagbanwa:                           {[2]rune{0x001760, 0x00177F}, "Tagbanwa"},
	BlockKhmer:                              {[2]rune{0x001780, 0x0017FF}, "Khmer"},
	BlockMongolian:                          {[2]rune{0x001800, 0x0018AF}, "Mongolian"},
	BlockUnifiedCanadianAboriginalSyllabicsExtended: {[2]rune{0x0018B0, 0x0018FF}, "Unified Canadian Aboriginal Syllabics Extended"},
	BlockLimbu:                               {[2]rune{0x001900, 0x00194F}, "Limbu"},
	BlockTaiLe:                               {[2]rune{0x001950, 0x00197F}, "Tai Le"},
	BlockNewTaiLue:                           {[2]rune{0x001980, 0x0019DF}, "New Tai Lue"},
	BlockKhmerSymbols:                        {[2]rune{0x0019E0, 0x0019FF}, "Khmer Symbols"},
	BlockBuginese:                            {[2]rune{0x001A00, 0x001A1F}, "Buginese"},
	BlockTaiTham:                             {[2]rune{0x001A20, 0x001AAF}, "Tai Tham"},
	BlockCombiningDiacriticalMarksExtended:   {[2]rune{0x001AB0, 0x001AFF}, "Combining Diacritical Marks Extended"},
	BlockBalinese:                            {[2]rune{0x001B00, 0x001B7F}, "Balinese"},
	BlockSundanese:                           {[2]rune{0x001B80, 0x001BBF}, "Sundanese"},
	BlockBatak:                               {[2]rune{0x001BC0, 0x001BFF}, "Batak"},
	BlockLepcha:                              {[2]rune{0x001C00, 0x001C4F}, "Lepcha"},
	BlockOlChiki:                             {[2]rune{0x001C50, 0x001C7F}, "Ol Chiki"},
	BlockCyrillicExtendedC:                   {[2]rune{0x001C80, 0x001C8F}, "Cyrillic Extended-C"},
	BlockGeorgianExtended:                    {[2]rune{0x001C90, 0x001CBF}, "Georgian Extended"},
	BlockSundaneseSupplement:                 {[2]rune{0x001CC0, 0x001CCF}, "Sundanese Supplement"},
	BlockVedicExtensions:                     {[2]rune{0x001CD0, 0x001CFF}, "Vedic Extensions"},
	BlockPhoneticExtensions:                  {[2]rune{0x001D00, 0x001D7F}, "Phonetic Extensions"},
	BlockPhoneticExtensionsSupplement:        {[2]rune{0x001D80, 0x001DBF}, "Phonetic Extensions Supplement"},
	BlockCombiningDiacriticalMarksSupplement: {[2]rune{0x001DC0, 0x001DFF}, "Combining Diacritical Marks Supplement"},
	BlockLatinExtendedAdditional:             {[2]rune{0x001E00, 0x001EFF}, "Latin Extended Additional"},
	BlockGreekExtended:                       {[2]rune{0x001F00, 0x001FFF}, "Greek Extended"},
	BlockGeneralPunctuation:                  {[2]rune{0x002000, 0x00206F}, "General Punctuation"},
	BlockSuperscriptsandSubscripts:           {[2]rune{0x002070, 0x00209F}, "Superscripts and Subscripts"},
	BlockCurrencySymbols:                     {[2]rune{0x0020A0, 0x0020CF}, "Currency Symbols"},
	BlockCombiningDiacriticalMarksforSymbols: {[2]rune{0x0020D0, 0x0020FF}, "Combining Diacritical Marks for Symbols"},
	BlockLetterlikeSymbols:                   {[2]rune{0x002100, 0x00214F}, "Letterlike Symbols"},
	BlockNumberForms:                         {[2]rune{0x002150, 0x00218F}, "Number Forms"},
	BlockArrows:                              {[2]rune{0x002190, 0x0021FF}, "Arrows"},
	BlockMathematicalOperators:               {[2]rune{0x002200, 0x0022FF}, "Mathematical Operators"},
	BlockMiscellaneousTechnical:              {[2]rune{0x002300, 0x0023FF}, "Miscellaneous Technical"},
	BlockControlPictures:                     {[2]rune{0x002400, 0x00243F}, "Control Pictures"},
	BlockOpticalCharacterRecognition:         {[2]rune{0x002440, 0x00245F}, "Optical Character Recognition"},
	BlockEnclosedAlphanumerics:               {[2]rune{0x002460, 0x0024FF}, "Enclosed Alphanumerics"},
	BlockBoxDrawing:                          {[2]rune{0x002500, 0x00257F}, "Box Drawing"},
	BlockBlockElements:                       {[2]rune{0x002580, 0x00259F}, "Block Elements"},
	BlockGeometricShapes:                     {[2]rune{0x0025A0, 0x0025FF}, "Geometric Shapes"},
	BlockMiscellaneousSymbols:                {[2]rune{0x002600, 0x0026FF}, "Miscellaneous Symbols"},
	BlockDingbats:                            {[2]rune{0x002700, 0x0027BF}, "Dingbats"},
	BlockMiscellaneousMathematicalSymbolsA:   {[2]rune{0x0027C0, 0x0027EF}, "Miscellaneous Mathematical Symbols-A"},
	BlockSupplementalArrowsA:                 {[2]rune{0x0027F0, 0x0027FF}, "Supplemental Arrows-A"},
	BlockBraillePatterns:                     {[2]rune{0x002800, 0x0028FF}, "Braille Patterns"},
	BlockSupplementalArrowsB:                 {[2]rune{0x002900, 0x00297F}, "Supplemental Arrows-B"},
	BlockMiscellaneousMathematicalSymbolsB:   {[2]rune{0x002980, 0x0029FF}, "Miscellaneous Mathematical Symbols-B"},
	BlockSupplementalMathematicalOperators:   {[2]rune{0x002A00, 0x002AFF}, "Supplemental Mathematical Operators"},
	BlockMiscellaneousSymbolsandArrows:       {[2]rune{0x002B00, 0x002BFF}, "Miscellaneous Symbols and Arrows"},
	BlockGlagolitic:                          {[2]rune{0x002C00, 0x002C5F}, "Glagolitic"},
	BlockLatinExtendedC:                      {[2]rune{0x002C60, 0x002C7F}, "Latin Extended-C"},
	BlockCoptic:                              {[2]rune{0x002C80, 0x002CFF}, "Coptic"},
	BlockGeorgianSupplement:                  {[2]rune{0x002D00, 0x002D2F}, "Georgian Supplement"},
	BlockTifinagh:                            {[2]rune{0x002D30, 0x002D7F}, "Tifinagh"},
	BlockEthiopicExtended:                    {[2]rune{0x002D80, 0x002DDF}, "Ethiopic Extended"},
	BlockCyrillicExtendedA:                   {[2]rune{0x002DE0, 0x002DFF}, "Cyrillic Extended-A"},
	BlockSupplementalPunctuation:             {[2]rune{0x002E00, 0x002E7F}, "Supplemental Punctuation"},
	BlockCJKRadicalsSupplement:               {[2]rune{0x002E80, 0x002EFF}, "CJK Radicals Supplement"},
	BlockKangxiRadicals:                      {[2]rune{0x002F00, 0x002FDF}, "Kangxi Radicals"},
	BlockIdeographicDescriptionCharacters:    {[2]rune{0x002FF0, 0x002FFF}, "Ideographic Description Characters"},
	BlockCJKSymbolsandPunctuation:            {[2]rune{0x003000, 0x00303F}, "CJK Symbols and Punctuation"},
	BlockHiragana:                            {[2]rune{0x003040, 0x00309F}, "Hiragana"},
	BlockKatakana:                            {[2]rune{0x0030A0, 0x0030FF}, "Katakana"},
	BlockBopomofo:                            {[2]rune{0x003100, 0x00312F}, "Bopomofo"},
	BlockHangulCompatibilityJamo:             {[2]rune{0x003130, 0x00318F}, "Hangul Compatibility Jamo"},
	BlockKanbun:                              {[2]rune{0x003190, 0x00319F}, "Kanbun"},
	BlockBopomofoExtended:                    {[2]rune{0x0031A0, 0x0031BF}, "Bopomofo Extended"},
	BlockCJKStrokes:                          {[2]rune{0x0031C0, 0x0031EF}, "CJK Strokes"},
	BlockKatakanaPhoneticExtensions:          {[2]rune{0x0031F0, 0x0031FF}, "Katakana Phonetic Extensions"},
	BlockEnclosedCJKLettersandMonths:         {[2]rune{0x003200, 0x0032FF}, "Enclosed CJK Letters and Months"},
	BlockCJKCompatibility:                    {[2]rune{0x003300, 0x0033FF}, "CJK Compatibility"},
	BlockCJKUnifiedIdeographsExtensionA:      {[2]rune{0x003400, 0x004DBF}, "CJK Unified Ideographs Extension A"},
	BlockYijingHexagramSymbols:               {[2]rune{0x004DC0, 0x004DFF}, "Yijing Hexagram Symbols"},
	BlockCJKUnifiedIdeographs:                {[2]rune{0x004E00, 0x009FFF}, "CJK Unified Ideographs"},
	BlockYiSyllables:                         {[2]rune{0x00A000, 0x00A48F}, "Yi Syllables"},
	BlockYiRadicals:                          {[2]rune{0x00A490, 0x00A4CF}, "Yi Radicals"},
	BlockLisu:                                {[2]rune{0x00A4D0, 0x00A4FF}, "Lisu"},
	BlockVai:                                 {[2]rune{0x00A500, 0x00A63F}, "Vai"},
	BlockCyrillicExtendedB:                   {[2]rune{0x00A640, 0x00A69F}, "Cyrillic Extended-B"},
	BlockBamum:                               {[2]rune{0x00A6A0, 0x00A6FF}, "Bamum"},
	BlockModifierToneLetters:                 {[2]rune{0x00A700, 0x00A71F}, "Modifier Tone Letters"},
	BlockLatinExtendedD:                      {[2]rune{0x00A720, 0x00A7FF}, "Latin Extended-D"},
	BlockSylotiNagri:                         {[2]rune{0x00A800, 0x00A82F}, "Syloti Nagri"},
	BlockCommonIndicNumberForms:              {[2]rune{0x00A830, 0x00A83F}, "Common Indic Number Forms"},
	BlockPhagspa:                             {[2]rune{0x00A840, 0x00A87F}, "Phags-pa"},
	BlockSaurashtra:                          {[2]rune{0x00A880, 0x00A8DF}, "Saurashtra"},
	BlockDevanagariExtended:                  {[2]rune{0x00A8E0, 0x00A8FF}, "Devanagari Extended"},
	BlockKayahLi:                             {[2]rune{0x00A900, 0x00A92F}, "Kayah Li"},
	BlockRejang:                              {[2]rune{0x00A930, 0x00A95F}, "Rejang"},
	BlockHangulJamoExtendedA:                 {[2]rune{0x00A960, 0x00A97F}, "Hangul Jamo Extended-A"},
	BlockJavanese:                            {[2]rune{0x00A980, 0x00A9DF}, "Javanese"},
	BlockMyanmarExtendedB:                    {[2]rune{0x00A9E0, 0x00A9FF}, "Myanmar Extended-B"},
	BlockCham:                                {[2]rune{0x00AA00, 0x00AA5F}, "Cham"},
	BlockMyanmarExtendedA:                    {[2]rune{0x00AA60, 0x00AA7F}, "Myanmar Extended-A"},
	BlockTaiViet:                             {[2]rune{0x00AA80, 0x00AADF}, "Tai Viet"},
	BlockMeeteiMayekExtensions:               {[2]rune{0x00AAE0, 0x00AAFF}, "Meetei Mayek Extensions"},
	BlockEthiopicExtendedA:                   {[2]rune{0x00AB00, 0x00AB2F}, "Ethiopic Extended-A"},
	BlockLatinExtendedE:                      {[2]rune{0x00AB30, 0x00AB6F}, "Latin Extended-E"},
	BlockCherokeeSupplement:                  {[2]rune{0x00AB70, 0x00ABBF}, "Cherokee Supplement"},
	BlockMeeteiMayek:                         {[2]rune{0x00ABC0, 0x00ABFF}, "Meetei Mayek"},
	BlockHangulSyllables:                     {[2]rune{0x00AC00, 0x00D7AF}, "Hangul Syllables"},
	BlockHangulJamoExtendedB:                 {[2]rune{0x00D7B0, 0x00D7FF}, "Hangul Jamo Extended-B"},
	BlockHighSurrogates:                      {[2]rune{0x00D800, 0x00DB7F}, "High Surrogates"},
	BlockHighPrivateUseSurrogates:            {[2]rune{0x00DB80, 0x00DBFF}, "High Private Use Surrogates"},
	BlockLowSurrogates:                       {[2]rune{0x00DC00, 0x00DFFF}, "Low Surrogates"},
	BlockPrivateUseArea:                      {[2]rune{0x00E000, 0x00F8FF}, "Private Use Area"},
	BlockCJKCompatibilityIdeographs:          {[2]rune{0x00F900, 0x00FAFF}, "CJK Compatibility Ideographs"},
	BlockAlphabeticPresentationForms:         {[2]rune{0x00FB00, 0x00FB4F}, "Alphabetic Presentation Forms"},
	BlockArabicPresentationFormsA:            {[2]rune{0x00FB50, 0x00FDFF}, "Arabic Presentation Forms-A"},
	BlockVariationSelectors:                  {[2]rune{0x00FE00, 0x00FE0F}, "Variation Selectors"},
	BlockVerticalForms:                       {[2]rune{0x00FE10, 0x00FE1F}, "Vertical Forms"},
	BlockCombiningHalfMarks:                  {[2]rune{0x00FE20, 0x00FE2F}, "Combining Half Marks"},
	BlockCJKCompatibilityForms:               {[2]rune{0x00FE30, 0x00FE4F}, "CJK Compatibility Forms"},
	BlockSmallFormVariants:                   {[2]rune{0x00FE50, 0x00FE6F}, "Small Form Variants"},
	BlockArabicPresentationFormsB:            {[2]rune{0x00FE70, 0x00FEFF}, "Arabic Presentation Forms-B"},
	BlockHalfwidthandFullwidthForms:          {[2]rune{0x00FF00, 0x00FFEF}, "Halfwidth and Fullwidth Forms"},
	BlockSpecials:                            {[2]rune{0x00FFF0, 0x00FFFF}, "Specials"},
	BlockLinearBSyllabary:                    {[2]rune{0x010000, 0x01007F}, "Linear B Syllabary"},
	BlockLinearBIdeograms:                    {[2]rune{0x010080, 0x0100FF}, "Linear B Ideograms"},
	BlockAegeanNumbers:                       {[2]rune{0x010100, 0x01013F}, "Aegean Numbers"},
	BlockAncientGreekNumbers:                 {[2]rune{0x010140, 0x01018F}, "Ancient Greek Numbers"},
	BlockAncientSymbols:                      {[2]rune{0x010190, 0x0101CF}, "Ancient Symbols"},
	BlockPhaistosDisc:                        {[2]rune{0x0101D0, 0x0101FF}, "Phaistos Disc"},
	BlockLycian:                              {[2]rune{0x010280, 0x01029F}, "Lycian"},
	BlockCarian:                              {[2]rune{0x0102A0, 0x0102DF}, "Carian"},
	BlockCopticEpactNumbers:                  {[2]rune{0x0102E0, 0x0102FF}, "Coptic Epact Numbers"},
	BlockOldItalic:                           {[2]rune{0x010300, 0x01032F}, "Old Italic"},
	BlockGothic:                              {[2]rune{0x010330, 0x01034F}, "Gothic"},
	BlockOldPermic:                           {[2]rune{0x010350, 0x01037F}, "Old Permic"},
	BlockUgaritic:                            {[2]rune{0x010380, 0x01039F}, "Ugaritic"},
	BlockOldPersian:                          {[2]rune{0x0103A0, 0x0103DF}, "Old Persian"},
	BlockDeseret:                             {[2]rune{0x010400, 0x01044F}, "Deseret"},
	BlockShavian:                             {[2]rune{0x010450, 0x01047F}, "Shavian"},
	BlockOsmanya:                             {[2]rune{0x010480, 0x0104AF}, "Osmanya"},
	BlockOsage:                               {[2]rune{0x0104B0, 0x0104FF}, "Osage"},
	BlockElbasan:                             {[2]rune{0x010500, 0x01052F}, "Elbasan"},
	BlockCaucasianAlbanian:                   {[2]rune{0x010530, 0x01056F}, "Caucasian Albanian"},
	BlockVithkuqi:                            {[2]rune{0x010570, 0x0105BF}, "Vithkuqi"},
	BlockLinearA:                             {[2]rune{0x010600, 0x01077F}, "Linear A"},
	BlockLatinExtendedF:                      {[2]rune{0x010780, 0x0107BF}, "Latin Extended-F"},
	BlockCypriotSyllabary:                    {[2]rune{0x010800, 0x01083F}, "Cypriot Syllabary"},
	BlockImperialAramaic:                     {[2]rune{0x010840, 0x01085F}, "Imperial Aramaic"},
	BlockPalmyrene:                           {[2]rune{0x010860, 0x01087F}, "Palmyrene"},
	BlockNabataean:                           {[2]rune{0x010880, 0x0108AF}, "Nabataean"},
	BlockHatran:                              {[2]rune{0x0108E0, 0x0108FF}, "Hatran"},
	BlockPhoenician:                          {[2]rune{0x010900, 0x01091F}, "Phoenician"},
	BlockLydian:                              {[2]rune{0x010920, 0x01093F}, "Lydian"},
	BlockMeroiticHieroglyphs:                 {[2]rune{0x010980, 0x01099F}, "Meroitic Hieroglyphs"},
	BlockMeroiticCursive:                     {[2]rune{0x0109A0, 0x0109FF}, "Meroitic Cursive"},
	BlockKharoshthi:                          {[2]rune{0x010A00, 0x010A5F}, "Kharoshthi"},
	BlockOldSouthArabian:                     {[2]rune{0x010A60, 0x010A7F}, "Old South Arabian"},
	BlockOldNorthArabian:                     {[2]rune{0x010A80, 0x010A9F}, "Old North Arabian"},
	BlockManichaean:                          {[2]rune{0x010AC0, 0x010AFF}, "Manichaean"},
	BlockAvestan:                             {[2]rune{0x010B00, 0x010B3F}, "Avestan"},
	BlockInscriptionalParthian:               {[2]rune{0x010B40, 0x010B5F}, "Inscriptional Parthian"},
	BlockInscriptionalPahlavi:                {[2]rune{0x010B60, 0x010B7F}, "Inscriptional Pahlavi"},
	BlockPsalterPahlavi:                      {[2]rune{0x010B80, 0x010BAF}, "Psalter Pahlavi"},
	BlockOldTurkic:                           {[2]rune{0x010C00, 0x010C4F}, "Old Turkic"},
	BlockOldHungarian:                        {[2]rune{0x010C80, 0x010CFF}, "Old Hungarian"},
	BlockHanifiRohingya:                      {[2]rune{0x010D00, 0x010D3F}, "Hanifi Rohingya"},
	BlockRumiNumeralSymbols:                  {[2]rune{0x010E60, 0x010E7F}, "Rumi Numeral Symbols"},
	BlockYezidi:                              {[2]rune{0x010E80, 0x010EBF}, "Yezidi"},
	BlockOldSogdian:                          {[2]rune{0x010F00, 0x010F2F}, "Old Sogdian"},
	BlockSogdian:                             {[2]rune{0x010F30, 0x010F6F}, "Sogdian"},
	BlockOldUyghur:                           {[2]rune{0x010F70, 0x010FAF}, "Old Uyghur"},
	BlockChorasmian:                          {[2]rune{0x010FB0, 0x010FDF}, "Chorasmian"},
	BlockElymaic:                             {[2]rune{0x010FE0, 0x010FFF}, "Elymaic"},
	BlockBrahmi:                              {[2]rune{0x011000, 0x01107F}, "Brahmi"},
	BlockKaithi:                              {[2]rune{0x011080, 0x0110CF}, "Kaithi"},
	BlockSoraSompeng:                         {[2]rune{0x0110D0, 0x0110FF}, "Sora Sompeng"},
	BlockChakma:                              {[2]rune{0x011100, 0x01114F}, "Chakma"},
	BlockMahajani:                            {[2]rune{0x011150, 0x01117F}, "Mahajani"},
	BlockSharada:                             {[2]rune{0x011180, 0x0111DF}, "Sharada"},
	BlockSinhalaArchaicNumbers:               {[2]rune{0x0111E0, 0x0111FF}, "Sinhala Archaic Numbers"},
	BlockKhojki:                              {[2]rune{0x011200, 0x01124F}, "Khojki"},
	BlockMultani:                             {[2]rune{0x011280, 0x0112AF}, "Multani"},
	BlockKhudawadi:                           {[2]rune{0x0112B0, 0x0112FF}, "Khudawadi"},
	BlockGrantha:                             {[2]rune{0x011300, 0x01137F}, "Grantha"},
	BlockNewa:                                {[2]rune{0x011400, 0x01147F}, "Newa"},
	BlockTirhuta:                             {[2]rune{0x011480, 0x0114DF}, "Tirhuta"},
	BlockSiddham:                             {[2]rune{0x011580, 0x0115FF}, "Siddham"},
	BlockModi:                                {[2]rune{0x011600, 0x01165F}, "Modi"},
	BlockMongolianSupplement:                 {[2]rune{0x011660, 0x01167F}, "Mongolian Supplement"},
	BlockTakri:                               {[2]rune{0x011680, 0x0116CF}, "Takri"},
	BlockAhom:                                {[2]rune{0x011700, 0x01174F}, "Ahom"},
	BlockDogra:                               {[2]rune{0x011800, 0x01184F}, "Dogra"},
	BlockWarangCiti:                          {[2]rune{0x0118A0, 0x0118FF}, "Warang Citi"},
	BlockDivesAkuru:                          {[2]rune{0x011900, 0x01195F}, "Dives Akuru"},
	BlockNandinagari:                         {[2]rune{0x0119A0, 0x0119FF}, "Nandinagari"},
	BlockZanabazarSquare:                     {[2]rune{0x011A00, 0x011A4F}, "Zanabazar Square"},
	BlockSoyombo:                             {[2]rune{0x011A50, 0x011AAF}, "Soyombo"},
	BlockUnifiedCanadianAboriginalSyllabicsExtendedA: {[2]rune{0x011AB0, 0x011ABF}, "Unified Canadian Aboriginal Syllabics Extended-A"},
	BlockPauCinHau:                            {[2]rune{0x011AC0, 0x011AFF}, "Pau Cin Hau"},
	BlockBhaiksuki:                            {[2]rune{0x011C00, 0x011C6F}, "Bhaiksuki"},
	BlockMarchen:                              {[2]rune{0x011C70, 0x011CBF}, "Marchen"},
	BlockMasaramGondi:                         {[2]rune{0x011D00, 0x011D5F}, "Masaram Gondi"},
	BlockGunjalaGondi:                         {[2]rune{0x011D60, 0x011DAF}, "Gunjala Gondi"},
	BlockMakasar:                              {[2]rune{0x011EE0, 0x011EFF}, "Makasar"},
	BlockLisuSupplement:                       {[2]rune{0x011FB0, 0x011FBF}, "Lisu Supplement"},
	BlockTamilSupplement:                      {[2]rune{0x011FC0, 0x011FFF}, "Tamil Supplement"},
	BlockCuneiform:                            {[2]rune{0x012000, 0x0123FF}, "Cuneiform"},
	BlockCuneiformNumbersandPunctuation:       {[2]rune{0x012400, 0x01247F}, "Cuneiform Numbers and Punctuation"},
	BlockEarlyDynasticCuneiform:               {[2]rune{0x012480, 0x01254F}, "Early Dynastic Cuneiform"},
	BlockCyproMinoan:                          {[2]rune{0x012F90, 0x012FFF}, "Cypro-Minoan"},
	BlockEgyptianHieroglyphs:                  {[2]rune{0x013000, 0x01342F}, "Egyptian Hieroglyphs"},
	BlockEgyptianHieroglyphFormatControls:     {[2]rune{0x013430, 0x01343F}, "Egyptian Hieroglyph Format Controls"},
	BlockAnatolianHieroglyphs:                 {[2]rune{0x014400, 0x01467F}, "Anatolian Hieroglyphs"},
	BlockBamumSupplement:                      {[2]rune{0x016800, 0x016A3F}, "Bamum Supplement"},
	BlockMro:                                  {[2]rune{0x016A40, 0x016A6F}, "Mro"},
	BlockTangsa:                               {[2]rune{0x016A70, 0x016ACF}, "Tangsa"},
	BlockBassaVah:                             {[2]rune{0x016AD0, 0x016AFF}, "Bassa Vah"},
	BlockPahawhHmong:                          {[2]rune{0x016B00, 0x016B8F}, "Pahawh Hmong"},
	BlockMedefaidrin:                          {[2]rune{0x016E40, 0x016E9F}, "Medefaidrin"},
	BlockMiao:                                 {[2]rune{0x016F00, 0x016F9F}, "Miao"},
	BlockIdeographicSymbolsandPunctuation:     {[2]rune{0x016FE0, 0x016FFF}, "Ideographic Symbols and Punctuation"},
	BlockTangut:                               {[2]rune{0x017000, 0x0187FF}, "Tangut"},
	BlockTangutComponents:                     {[2]rune{0x018800, 0x018AFF}, "Tangut Components"},
	BlockKhitanSmallScript:                    {[2]rune{0x018B00, 0x018CFF}, "Khitan Small Script"},
	BlockTangutSupplement:                     {[2]rune{0x018D00, 0x018D7F}, "Tangut Supplement"},
	BlockKanaExtendedB:                        {[2]rune{0x01AFF0, 0x01AFFF}, "Kana Extended-B"},
	BlockKanaSupplement:                       {[2]rune{0x01B000, 0x01B0FF}, "Kana Supplement"},
	BlockKanaExtendedA:                        {[2]rune{0x01B100, 0x01B12F}, "Kana Extended-A"},
	BlockSmallKanaExtension:                   {[2]rune{0x01B130, 0x01B16F}, "Small Kana Extension"},
	BlockNushu:                                {[2]rune{0x01B170, 0x01B2FF}, "Nushu"},
	BlockDuployan:                             {[2]rune{0x01BC00, 0x01BC9F}, "Duployan"},
	BlockShorthandFormatControls:              {[2]rune{0x01BCA0, 0x01BCAF}, "Shorthand Format Controls"},
	BlockZnamennyMusicalNotation:              {[2]rune{0x01CF00, 0x01CFCF}, "Znamenny Musical Notation"},
	BlockByzantineMusicalSymbols:              {[2]rune{0x01D000, 0x01D0FF}, "Byzantine Musical Symbols"},
	BlockMusicalSymbols:                       {[2]rune{0x01D100, 0x01D1FF}, "Musical Symbols"},
	BlockAncientGreekMusicalNotation:          {[2]rune{0x01D200, 0x01D24F}, "Ancient Greek Musical Notation"},
	BlockMayanNumerals:                        {[2]rune{0x01D2E0, 0x01D2FF}, "Mayan Numerals"},
	BlockTaiXuanJingSymbols:                   {[2]rune{0x01D300, 0x01D35F}, "Tai Xuan Jing Symbols"},
	BlockCountingRodNumerals:                  {[2]rune{0x01D360, 0x01D37F}, "Counting Rod Numerals"},
	BlockMathematicalAlphanumericSymbols:      {[2]rune{0x01D400, 0x01D7FF}, "Mathematical Alphanumeric Symbols"},
	BlockSuttonSignWriting:                    {[2]rune{0x01D800, 0x01DAAF}, "Sutton SignWriting"},
	BlockLatinExtendedG:                       {[2]rune{0x01DF00, 0x01DFFF}, "Latin Extended-G"},
	BlockGlagoliticSupplement:                 {[2]rune{0x01E000, 0x01E02F}, "Glagolitic Supplement"},
	BlockNyiakengPuachueHmong:                 {[2]rune{0x01E100, 0x01E14F}, "Nyiakeng Puachue Hmong"},
	BlockToto:                                 {[2]rune{0x01E290, 0x01E2BF}, "Toto"},
	BlockWancho:                               {[2]rune{0x01E2C0, 0x01E2FF}, "Wancho"},
	BlockEthiopicExtendedB:                    {[2]rune{0x01E7E0, 0x01E7FF}, "Ethiopic Extended-B"},
	BlockMendeKikakui:                         {[2]rune{0x01E800, 0x01E8DF}, "Mende Kikakui"},
	BlockAdlam:                                {[2]rune{0x01E900, 0x01E95F}, "Adlam"},
	BlockIndicSiyaqNumbers:                    {[2]rune{0x01EC70, 0x01ECBF}, "Indic Siyaq Numbers"},
	BlockOttomanSiyaqNumbers:                  {[2]rune{0x01ED00, 0x01ED4F}, "Ottoman Siyaq Numbers"},
	BlockArabicMathematicalAlphabeticSymbols:  {[2]rune{0x01EE00, 0x01EEFF}, "Arabic Mathematical Alphabetic Symbols"},
	BlockMahjongTiles:                         {[2]rune{0x01F000, 0x01F02F}, "Mahjong Tiles"},
	BlockDominoTiles:                          {[2]rune{0x01F030, 0x01F09F}, "Domino Tiles"},
	BlockPlayingCards:                         {[2]rune{0x01F0A0, 0x01F0FF}, "Playing Cards"},
	BlockEnclosedAlphanumericSupplement:       {[2]rune{0x01F100, 0x01F1FF}, "Enclosed Alphanumeric Supplement"},
	BlockEnclosedIdeographicSupplement:        {[2]rune{0x01F200, 0x01F2FF}, "Enclosed Ideographic Supplement"},
	BlockMiscellaneousSymbolsandPictographs:   {[2]rune{0x01F300, 0x01F5FF}, "Miscellaneous Symbols and Pictographs"},
	BlockEmoticons:                            {[2]rune{0x01F600, 0x01F64F}, "Emoticons"},
	BlockOrnamentalDingbats:                   {[2]rune{0x01F650, 0x01F67F}, "Ornamental Dingbats"},
	BlockTransportandMapSymbols:               {[2]rune{0x01F680, 0x01F6FF}, "Transport and Map Symbols"},
	BlockAlchemicalSymbols:                    {[2]rune{0x01F700, 0x01F77F}, "Alchemical Symbols"},
	BlockGeometricShapesExtended:              {[2]rune{0x01F780, 0x01F7FF}, "Geometric Shapes Extended"},
	BlockSupplementalArrowsC:                  {[2]rune{0x01F800, 0x01F8FF}, "Supplemental Arrows-C"},
	BlockSupplementalSymbolsandPictographs:    {[2]rune{0x01F900, 0x01F9FF}, "Supplemental Symbols and Pictographs"},
	BlockChessSymbols:                         {[2]rune{0x01FA00, 0x01FA6F}, "Chess Symbols"},
	BlockSymbolsandPictographsExtendedA:       {[2]rune{0x01FA70, 0x01FAFF}, "Symbols and Pictographs Extended-A"},
	BlockSymbolsforLegacyComputing:            {[2]rune{0x01FB00, 0x01FBFF}, "Symbols for Legacy Computing"},
	BlockCJKUnifiedIdeographsExtensionB:       {[2]rune{0x020000, 0x02A6DF}, "CJK Unified Ideographs Extension B"},
	BlockCJKUnifiedIdeographsExtensionC:       {[2]rune{0x02A700, 0x02B73F}, "CJK Unified Ideographs Extension C"},
	BlockCJKUnifiedIdeographsExtensionD:       {[2]rune{0x02B740, 0x02B81F}, "CJK Unified Ideographs Extension D"},
	BlockCJKUnifiedIdeographsExtensionE:       {[2]rune{0x02B820, 0x02CEAF}, "CJK Unified Ideographs Extension E"},
	BlockCJKUnifiedIdeographsExtensionF:       {[2]rune{0x02CEB0, 0x02EBEF}, "CJK Unified Ideographs Extension F"},
	BlockCJKCompatibilityIdeographsSupplement: {[2]rune{0x02F800, 0x02FA1F}, "CJK Compatibility Ideographs Supplement"},
	BlockCJKUnifiedIdeographsExtensionG:       {[2]rune{0x030000, 0x03134F}, "CJK Unified Ideographs Extension G"},
	BlockTags:                                 {[2]rune{0x0E0000, 0x0E007F}, "Tags"},
	BlockVariationSelectorsSupplement:         {[2]rune{0x0E0100, 0x0E01EF}, "Variation Selectors Supplement"},
	BlockSupplementaryPrivateUseAreaA:         {[2]rune{0x0F0000, 0x0FFFFF}, "Supplementary Private Use Area-A"},
	BlockSupplementaryPrivateUseAreaB:         {[2]rune{0x100000, 0x10FFFF}, "Supplementary Private Use Area-B"},
}
