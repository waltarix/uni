BEGIN {
    FS                    = ";"
    PROCINFO["sorted_in"] = "@ind_num_asc"
    loadwidths()
}

{
    codepoint = strtonum("0x" $1)
    name      = $2
    cat       = $3

    if (index(name, "<") == 1) {
        # Control characters all have the name as <control>, which isn't very
        # useful. The "obsolete" Unicode 1 name field is more useful.
        if (length($11) > 1)
            name = $11
        # Record ranges:
        #   3400;<CJK Ideograph Extension A, First>;Lo;0;L;;;;;N;;;;;
        #   4DBF;<CJK Ideograph Extension A, Last>;Lo;0;L;;;;;N;;;;;
        else if (match(name, "First|Last>$") > 0) {
            x = gensub(", (First|Last)", "", 1, name)
            ranges[x] = sprintf("%s0x%06X, ", ranges[x], codepoint)
        }
    }

    codepoints = codepoints sprintf("\t0x%X: {0x%X, %s, Cat%s, \"%s\"},\n",
        codepoint, codepoint, widths[codepoint], cat, name)
}

END {
    print("// Code generated by gen.zsh; DO NOT EDIT\n\npackage unidata\n")

    print("// Codepoints that aren't listed individually.\n" \
          "var codepointRanges = []struct {\n" \
              "\trng  [2]rune\n" \
              "\tname string\n" \
          "}{")
    for (k in ranges) printf("\t{[2]rune{%s}, \"%s\"},\n", ranges[k], k)
    print("}\n")

    print("var Codepoints = map[rune]Codepoint{\n" codepoints "\n}\n")

    print("var htmlEntities = map[rune]string{")
    while (getline line <".cache/entities.json" > 0) {
        split(line, fields, /["\[\]]/)
        ent = fields[2]
        cp  = strtonum(fields[6])
        if (ent == "" || index(ent, ";") == 0)  # Don't include legacy w/o ;
            continue
        ent = gensub("^&|;$", "", "g", ent)

        # Prefer entities without capitals and shorter entities.
        # TODO: include all entities as a []string; select the first one with
        # "%(html)" and all of them with "%(htmls)", "%(html_all)" or something.
        if (!(cp in all) || (match(all[cp], /^[A-Z]/) && !match(ent, /^[A-Z]/)) || length(all[cp]) > length(ent))
            all[cp] = ent
    }
    for (k in all) printf("\t0x%02x: \"%s\",\n", k, all[k])
    print("}\n")

    print("var keysyms = map[rune]string{")
    while (getline line <".cache/keysymdef.h" > 0) {
        if (match(line, "^#define XK") == 0)
            continue
        split(line, fields, " ")
        all[strtonum(fields[3])] = gensub("^XK_", "", 1, fields[2])
    }
    for (k in all) printf("\t0x%02x: \"%s\",\n", k, all[k])
    print("}\n")

    print("var digraphs = map[rune]string{")
    while (getline line <".cache/rfc1345.txt" > 0) {
		if (index(line, "ISO-IR-") > 0)
            continue
        if (match(line, "^ .*?   +[0-9a-f]{4}") > 0) {
            split(line, fields, " ")
            all[strtonum("0x" fields[2])] = gensub("\"", "\\\\\"", "g", fields[1])
        }
    }
    all[0x00]   = "NU" # Correct for inconsistent line
    all[0x20ac] = "=e" # € (Euro)
    all[0x20bd] = "=R" # ₽ (Ruble); also =P and the only one with more than one digraph :-/
    for (k in all) printf("\t0x%02x: \"%s\",\n", k, all[k])
    print("}")
}

function loadwidths(      fields, width, cp, start, end, i) {
    while (getline line <".cache/EastAsianWidth.txt" > 0) {
        if (match(line, "^$|^#") > 0)
            continue

        split(line, fields, "[; ]+")
        split(fields[1], cp, /\.\./)
        start = strtonum("0x" cp[1])
        end   = strtonum("0x" (length(cp) > 1 ? cp[2] : cp[1]))

        switch (fields[2]) {
        case "A":  width = "WidthAmbiguous"; break
        case "F":  width = "WidthFullWidth"; break
        case "H":  width = "WidthHalfWidth"; break
        case "N":  width = "WidthNarrow";    break
        case "Na": width = "WidthNeutral";   break
        case "W":  width = "WidthWide";      break
        default:
            print("unknown width" width)
            exit 1
        }

        for (i=start; i<=end; i++)
            widths[i] = width
    }
}
