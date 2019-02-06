package GET;

use 5.16.0;
use strict;
use warnings;

use Moose;

use LWP::UserAgent ();

has 'User-Agent'   => ( is => 'ro', isa => 'Str' );
has 'Content-Type' => ( is => 'ro', isa => 'Str' );
has 'postfix'      => ( is => 'ro', isa => 'Str' );
has 'host'         => ( is => 'ro', isa => 'HashArray' );
has 'url'          => ( is => 'ro', isa => 'HashArray' );

sub prepare {
    my $self = shift;

    my $ua = LWP::UserAgent->new;

    $self->{host} = 'http://127.0.0.1:8080';

    $self->{url} = $self->{host} . $self->{postfix};

    $ua;
}

1;
# -------------------------------------------------------------------------------------------------
# -------------------------------------------------------------------------------------------------
# -------------------------------------------------------------------------------------------------
use 5.16.0;
use strict;
use warnings;

use DDP;

use HTTP::Response;
use Encode;

use LWP::UserAgent ();
use JSON::XS;
use utf8;

use Test::Deep;
use Test::Deep::JSON;

use lib '.';

# use get_3::MakeDump;

# MakeDump::make_dump();
my @answers = MakeDump::read_dump();

open( my $fh_ammo, "<:encoding(utf8)", "../ammo/phase_3_get.ammo" ) or die $!;

my $req = [];

my $cnt = 0;

use GET;

my $get = GET->new();

while (<$fh_ammo>) {
    if ( $_ !~ /^\s*$/ ) {

        if (/GET\s([^\s]+)/) {
            $get->{'postfix'} = $1 . '/';
        }
    }

    if ( $get->{'postfix'} ) {
        my $ua = $get->prepare();

        if (defined $get->{postfix}
            && $get->{postfix} !~ /recommend/
            && $get->{postfix} !~ /suggest/
            
            && $get->{postfix} !~ /interest/
            && $get->{postfix} !~ /likes/
            )
        {
            my $url = $get->{url};
            
            chop $url;

            my $response = $ua->get($url);

            my ( $code, $code_name ) = split /\s+/, $response->status_line();

            if ( 0+$code != 0+$answers[$cnt]->[2] ) {
                p $url;

                print "Expect: ";
                p $answers[$cnt]->[2];
                
                print "Got: ";
                p $code;

                p $answers[$cnt]->[3];

                say "cnt: $cnt ";

                exit(0);
            }

            if ( $response->is_success ) {
                my $message = $response->decoded_content;

                # p $message;
                $message = decode( 'utf-8', $message );
                $message = encode( 'utf-8', $message );

                my ( $method, $json ) = split /\n/, $message;

                my $json_ = JSON::XS->new->utf8->decode($json);

                if ( !cmp_deeply( $answers[$cnt]->[3], $json_ ) ) {
                    p $answers[$cnt]->[3];

                    say "Got:";
                    p $json_;

                    say $url;
                    say "cnt: $cnt ", $code, ' ', $answers[$cnt]->[2];

                    exit(0);
                }
            }

            # say "cnt: $cnt ", $code, ' ', $answers[$cnt]->[2];
        }

        $cnt++;

        $get = GET->new();
    }
}

say "\nFINISHED\n";
