BEGIN        { FS = " *[;#] *" }
/^$/ || /^#/ { next }

/^gc/ {
    short = $2
    long  = $3
    incl  = index($NF, "|") > 0 ? $NF : ""
    const = "Cat" gensub(/_/, "", "g", long)

    longConstants  = longConstants sprintf("\t%s // %-4s", const, short)
    if (length(incl) == 0)
        incl = "nil"
    else {
        longConstants = longConstants "(" incl ")"
        gsub(/ \| /, ", Cat", incl)
        incl = "[]Category{Cat" incl "}"
    }
    longConstants = longConstants "\n"

    shortConstants = shortConstants "\tCat" short " = " const "\n"
    categories     = categories sprintf("\t%s: {\"%s\", \"%s\", %s},\n", const, short, long, incl)
}

END {
    print("// Code generated by gen.zsh; DO NOT EDIT\n\npackage unidata\n")

    print("// Unicode general categories (long names).\nconst (\n" \
          "\tCatUnknown = Category(iota)\n" \
          longConstants ")\n")

    print("// Unicode general categories (short names).\n" \
          "const (\n" shortConstants ")\n")

    print("// Categories is a list of all categories.\n" \
          "var Categories = map[Category]struct {\n" \
              "\tShortName, Name string\n" \
              "\tInclude         []Category\n" \
          "}{\n" categories "}")
}
