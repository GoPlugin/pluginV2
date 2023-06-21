#!/usr/bin/perl

open(FH, "../onlygoFile.txt");
my @arr = <FH>;
close FH;

chomp @arr;

foreach my $tmp(@arr){

	`perl -i -p -e "s/smartcontractkit\/chainlink/pluginV2/g" $tmp`;
}
