#!/usr/bin/env bash

both/both > both.dot; neato -Tsvg both.dot > both.svg

pubmedarticle2dot/pubmedarticle2dot > pubmedarticle2dot.dot; neato -Tsvg pubmedarticle2dot.dot > pubmedarticle2dot.svg

MeSH2dot/MeSH2dot > MeSH2dot.dot; neato -Tsvg MeSH2dot.dot > MeSH2dot.svg

