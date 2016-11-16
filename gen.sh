#!/bin/sh

for OP in ItemSearch BrowseNodeLookup ItemLookup SimilarityLookup CartAdd CartClear CartCreate CartGet CartModify; do
  FN=$(ruby -ractive_support -e "print ActiveSupport::Inflector.underscore('$OP')")
  erb operation=$OP main.erb | gofmt > amazon/$FN.go
done
