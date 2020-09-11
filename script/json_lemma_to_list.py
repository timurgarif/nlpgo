# Convert Spacy json lemmata to a TSV(lemma, POS)

import re
import json
import argparse

parser = argparse.ArgumentParser(
    description="Convert Spacy json lemmata to a TSV(lemma, POS)."
)
parser.add_argument(
    "-src", dest="src", help="Path to a JSON file", required=True,
)

args = parser.parse_args()

with open(args.src, encoding="utf-8") as rf:
    dest = re.sub(r"\.*$", ".tsv", args.src)
    with open(dest, "w", buffering=20 * 1024, encoding="utf-8") as wf:
        list = []
        data = json.load(rf)
        for pos in data:
            upperPos = pos.upper()
            for lemma in data[pos]:
                wf.write(lemma + "\t" + upperPos + "\n")

