use 5.16.0;
use strict;
use warnings;

use DDP;
use utf8;

# use DBD::mysql;
use DBI;

use JSON::XS;

	binmode STDIN,  ":encoding(utf8)";
binmode STDOUT, ":encoding(utf8)";

sub connect_db {
    my ( $host, $login, $pass, $db, $table ) = @_;

    my $dbh = DBI->connect( "DBI:mysql:database=$db;host=$host",
        $login, $pass, { 'RaiseError' => 1 } )
        or die $!;

    $dbh->{mysql_enable_utf8} = 1;

    $dbh->do('set names utf8');

    $dbh;
}

sub empty_table {
    my $dbh = shift;

    $dbh->do("DELETE FROM accounts") or die $dbh->errstr();
}

my $dbh = connect_db( 'localhost', 'alexfmsu', '321678', 'highload2018',
    'accounts' );

empty_table($dbh);

my $json;

my ( @fields, @values );
my ( $fields, $values );

my $cnt = 0;

my @field_names
    = qw(id email fname sname phone sex birth country city joined status premium);

for my $filename ( 1 .. 3 ) {
    open( my $fh, "<:encoding(utf8)", 'data/accounts_' . $filename . '.json' )
        or die $!;

    $json = JSON::XS->new->utf8->decode(<$fh>);

    for ( $json->{'accounts'} ) {
        for my $arr (@$_) {
            @fields = ();
            @values = ();
                	
            for (@field_names) {
                if ( exists $arr->{$_} ) {
                    if($_ eq 'premium'){
                    	push @fields, 'premium_start';
	                    push @values, "'" . $arr->{$_}->{start} . "'";
	                
	                	push @fields, 'premium_finish';
	                    push @values, "'" . $arr->{$_}->{finish} . "'";
                    }else{
                    	push @fields, $_;
                    	push @values, "'" . $arr->{$_} . "'";
                	}
                }
            }

            $fields = join ',', @fields;
            $values = join ',', @values;

            my $req
                = 'INSERT INTO accounts('
                . $fields
                . ') values('
                . $values . ')';

            # p $req;
            # sleep(1);
            # say $req;

            $dbh->do($req) or die $dbh->errstr();

            # -----------------------------------
            if ( exists $arr->{likes} ) {
                my $req_l = 'INSERT INTO likes(id, id1, id2, ts) values';

                for ( @{ $arr->{likes} } ) {
                    $req_l
                        .= '(' . '0' . ',' . "'"
                        . $arr->{id} . "'," . "'"
                        . $_->{id} . "'," . "'"
                        . $_->{ts} . "'" . '),';
                }

                chop($req_l);
                $dbh->do($req_l) or die $dbh->errstr();
            }

            # -----------------------------------
            if ( exists $arr->{interests} ) {
                my $reqi = 'INSERT INTO interests(id, id1, interest) values';

                for ( @{ $arr->{interests} } ) {
                    $reqi
                        .= '(' . '0' . ',' . "'"
                        . $arr->{id} . "'," . "'"
                        . $_ . "'" . '),';
                }

                chop($reqi);

                # say $reqi;

                $dbh->do($reqi) or die $dbh->errstr();
            }

            say $cnt++;
        }
    }
}

$dbh->disconnect();

