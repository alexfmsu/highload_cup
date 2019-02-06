package POST;

use 5.16.0;
use strict;
use warnings;
use Moose;
use LWP::UserAgent ();

has 'User-Agent'   => ( is => 'ro', isa => 'Str' );
has 'Content-Type' => ( is => 'ro', isa => 'Str' );
has 'postfix'      => ( is => 'ro', isa => 'Str' );
has 'host'         => ( is => 'ro', isa => 'HashArray' );
has 'postfix'      => ( is => 'ro', isa => 'HashArray' );
has 'url'          => ( is => 'ro', isa => 'HashArray' );
has 'json'         => ( is => 'ro', isa => 'HashRef' );

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

use 5.16.0;
use strict;
use warnings;
use DDP;
use HTTP::Response;

use LWP::UserAgent ();
use JSON::XS qw(encode_json);

open( my $fh_answ, "<", "answers/phase_2_post.answ" ) or die $!;

my @post_answ;

my $ans_cnt = 0;

for (<$fh_answ>) {
    my ( $method, $url, $code, $json ) = split /\s+/, $_;

    push @post_answ, [ $method, $url, $code, $json ];

    # last if $ans_cnt++ > 5;
}

# p @post_answ;
# exit(0);

open( my $fh_ammo, "<", "ammo/phase_2_post.ammo" ) or die $!;

my $req = [];

my $cnt = 0;

use POST;

my $post = POST->new();

for (<$fh_ammo>) {
    if ( $_ !~ /^\s*$/ ) {

        # push @$req, $_;

        # sleep(1);
        # say $_;

        if (/^User-Agent:\s(.+)$/) {
            $post->{'User-Agent'} = $1;
        }

        if (/^Content-Type:\s(.+)$/) {
            $post->{'Content-Type:'} = $1;
        }

        if (/POST:(.+)$/) {
            $post->{'postfix'} = $1;

            # say $1;

            # exit(0);
        }elsif(/POST\s\/accounts\/(\d+)\//){
            $post->{'postfix'} = '/accounts/'.$1.'/';
        }

        if (/^\{(.+)\}\s*$/) {
            my $json = $1;

            # $json=~s/\:/=>/g;

            $post->{'json'} = $_;

            # p $json;
            # exit(1);
        }
    }

    if ( $post->{'json'} ) {
        my $ua = $post->prepare();

        # my $response = $ua->get('http://search.cpan.org/');
        # my $url = 'http://127.0.0.1:8080/accounts/new';

        if ( defined $post->{postfix}
            && $post->{postfix} eq '/accounts/new/' || $post->{postfix} =~ '/accounts/\d+/' )
        {
            my $url = $post->{url};
            chop $url;

            # say $url;
            # say $post->{json};
            # say 'http://127.0.0.1:8080/accounts/new';
            my $response = $ua->post( $url, Content => $post->{json} );

            my ( $code, $code_name ) = split /\s+/, $response->status_line();

            if ( $code ne $post_answ[$cnt]->[2] ) {
            	say $response->status_line();
                
                say "cnt: $cnt ", $code, ' ', $post_answ[$cnt]->[2];

                p $post;

                # say $post->{url};
                # say $post->{json};
                # last;
            }


            last if $cnt > 10;

        }
        
        $cnt++;
        
        $post = POST->new();

# my $response = $ua->post( 'http://127.0.0.1:8080/accounts/new', Content => JSON::XS->new->encode($json) );
# p $post->{json};

        # my $aa = JSON::XS->new->encode($post->{json});
        # say "JSON:";
        # p $aa;
        # if ( $response->is_success() ) {
        #     say $response->status_line();
        # }
        # else {
        #     print( "ERROR: " . $response->status_line() );
        # }

        # my $response = $ua->put(
        # 	'http://127.0.0.1:8080/accounts/new',
        # 	Content=>{'phone'=>'12', 'email'=>'sada'}
        # );

        # if ( $response->is_success ) {
        #     print $response->decoded_content;    # or whatever
        # }
        # else {
        # 	p $response;
        # 	say "err";
        # 	die;
        #     # die $response->status_line;
        # }

        # p @$req;
        # last;
    }
}

# p @$req;

# my $req = HTTP::Request->new( POST => $url );
# $req->content_type('application/json');
# my $json = {
#     "email"   => 'ssador@yahoo.com',
#     "fname"   => "Полина",
#     "sname"   => "Хопетачан",
#     "country" => "Голция",
#     "city"    => "Голция",
#     "joined"  => "Голция",
#     "birth"   => 736598811,
#     "id"      => 50000,
#     "sex"     => "f",
#     "phone"   => "Пdолина",
# };

# $req->content($json);

# my $ua  = LWP::UserAgent->new;    # You might want some options here
# my $res = $ua->request($req);

# my $ua = new LWP::UserAgent();

# my $js = ;
# say ('http://127.0.0.1:8080'.$post->{postfix});
# my $urll = 'http://127.0.0.1:8080/accounts/new';
# my $urll = 'http://127.0.0.1:8080'.$post->{postfix};
# say $post->{json};

# my $response = $ua->post( $post->{url}, Content => $post->{json} );
# p $post;
# exit(0);
# p $post;
# sleep 1;
# if ( defined $post->{preffix}  ) {

say "FINISHED";