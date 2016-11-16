#!/bin/sh

for OP in ItemSearch BrowseNodeLookup ItemLookup SimilarityLookup CartAdd CartClear CartCreate CartGet CartModify; do
  FN=$(ruby -ractive_support -e "print ActiveSupport::Inflector.underscore('$OP')")
  erb operation=$OP templates/main.erb | gofmt > amazon/$FN.go
  erb operation=$OP templates/test.erb | gofmt > amazon/${FN}_test.go
done
