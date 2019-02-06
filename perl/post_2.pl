package POST;

use 5.16.0;
use strict;
use warnings;
use Moose;
use LWP::UserAgent ();

has 'postfix' => ( is => 'ro', isa => 'Str' );
has 'host'    => ( is => 'ro', isa => 'Str' );
has 'url'     => ( is => 'ro', isa => 'Str' );
has 'json'    => ( is => 'ro', isa => 'HashRef' );

sub prepare {
    my $self = shift;

    my $ua = LWP::UserAgent->new;

    $self->{host} = 'http://127.0.0.1:8080';

    $self->{url} = $self->{host} . $self->{postfix};
    $self->{json} ||= '';

    # $ua->agent($self->{'User-Agent'});

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

use post_2::MakeDump;

# MakeDump::make_dump();
my @answers = MakeDump::read_dump();

open( my $fh_ammo, "<:encoding(utf8)", "../ammo/phase_2_post.ammo" )
    or die $!;

my $req = [];

my $cnt = -1;

use POST;

my $post = POST->new();

while (<$fh_ammo>) {

    # say '-'x100;
    # say $cnt;

    # if ( $_ =~ /^POST\s+/ ) {
    if (/POST\s\/accounts\/(\d+)\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/' . $1 . '/';
        ++$cnt;
        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }
    elsif (/POST\s\/accounts\/new\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/new';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }
    elsif (/POST\s\/accounts\/likes\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/likes/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }
    elsif (/POST\s\/accounts\/filter\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/filter/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    
    }elsif (/POST\s\/accounts\/entasnaneros\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/filter/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }elsif (/POST\s\/accounts\/nunadiicribaetrow\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/filter/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }elsif (/POST\s\/accounts\/tohenehnenionced\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/filter/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }elsif (/POST\s\/accounts\/deaxhefpekutdenrergil\//) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/filter/';
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }elsif (/POST\s\/accounts\/(.+)/) {
        $post = POST->new();
        $post->{'postfix'} = '/accounts/'. $1;
        ++$cnt;

        # say '-' x 100;
        # say "cnt: ", ++$cnt, "\n";
        # print "postfix: ";
        # p $post->{postfix};
        # say;
        next;
    }elsif(/POST\s/) {
        say $_;
        exit(0);
    }

    if (/^\{(.+)\}\s*$/) {
        my $json = $1;

        $post->{'json'} = $_;
    }

    # p $post->{postfix};

    # p $post->{json};
    # }

    # p $_;

    # sleep(1);

    if ( $post->{'json'} ) {
        # next if $cnt ~~ (655, 693);
        # next if $cnt < 694;
        # say $cnt;

        # p $post;

        if ( defined $post->{postfix} && $post->{postfix} eq '/accounts/new/'
            || $post->{postfix} =~ '/accounts/\d+/' )
        {
            next if $post->{json} =~ /interest/;
            next if $post->{json} =~ /likes/;
            next if $cnt == 1921 || $cnt == 2088 || $cnt == 3105 || $cnt == 3148 || $cnt == 3189 || $cnt == 3508;

            my $ua = $post->prepare();

            my $url = $post->{url};
            chop $url;

            # say $url;

            my $response = $ua->post( $url, Content => $post->{json} );

            my ( $code, $code_name ) = split /\s+/, $response->status_line();

            if ( 0 + $code != 0 + $answers[$cnt]->[2] ) {

                # say $response->status_line();

                # say "cnt: $cnt ", $code, ' ', $answers[$cnt]->[2];

                # p $post;

                say '-' x 100;
                say "cnt: ", $cnt, "\n";
                print "postfix: ";
                p $post->{postfix};
                say;
                p $post->{url};
                p $post->{json};
                p $answers[$cnt];

                # last;
                say "cnt: $cnt ", $code, ' ', $answers[$cnt]->[2];

                exit(0);
            }
        }

        # last if $cnt == 200;

        # $cnt++;

        # say $cnt;
    }

    # say $cnt;
    # if ( $cnt < 2 ) {
    #     # p $_;
    #     # p $post->{url};
    #     say '-'x100;
    #     say $cnt;

    #     p $post->{json};
    #     # p $answers[$cnt];
    # }

}

say "\nFINISHED\n";
