package MakeDump;

use 5.16.0;
use strict;
use warnings;

use DDP;

use JSON::XS;

use utf8;

binmode STDIN,  ":encoding(utf8)";
binmode STDOUT, ":encoding(utf8)";

use Exporter 'import';

our @EXPORT = qw(make_dump read_dump);

# our @EXPORT_OK = qw(functionB);

sub make_dump {
    open( my $fh_answ, "<:encoding(utf8)", "../answers/phase_2_post.answ" )
        or die $!;

    my @get_answ;

    my $cnt = 0;

    while (<$fh_answ>) {

        # next if /^\s*$/;

        my ( $method, $url, $code, $json ) = split /\t/, $_;

        if ( defined($json) && $json !~ /^\s*$/ ) {
            $json = JSON::XS->new->utf8->decode($json);
        }

        push @get_answ, [ $method, $url, $code, $json ];

        say $cnt++;

        # last if $ans_cnt++ > 10;
    }

    open my $fh, ">", "post_2_answers.json";
    print $fh JSON::XS->new->utf8->encode( \@get_answ );
    close $fh;
}

sub read_dump {
    open my $fh, "<", "post_2_answers.json";
    my $data = <$fh>;
    close $fh;

    $data = JSON::XS->new->utf8->decode($data );

    return @$data;
}