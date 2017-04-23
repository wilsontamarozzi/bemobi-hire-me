pg_hashids, generate short unique ids from integers
=========================================

Hashids is a small open-source library that generates short, unique, non-sequential ids from numbers.
It converts numbers like 347 into strings like “yr8”.
You can also decode those ids back. This is useful in bundling several parameters into one or simply using them as short UIDs.

You can use hashids to hide primary keys in your database. I've used the extension on several production databases. 

Tested PostgreSQL versions : 9.5.X and 9.6.X (Should work on older versions, just not tested)

It's using [hashids.c](https://github.com/tzvetkoff/hashids.c) under the hood. More information about hashids and it's implementations here: [hashids.org](http://hashids.org)

Installation
============

Make sure you have development packages installed for postgres and
build tools in general.

    make; sudo make install

Then in a psql session issue:

    CREATE extension pg_hashids;


Update
============

Install as usual.

    make; sudo make install
    
Then in a psql session issue:

    ALTER EXTENSION pg_hashids UPDATE;
    
or

    DROP EXTENSION pg_hashids;
    CREATE EXTENSION pg_hashids;
    
Check

    SELECT default_version, installed_version FROM pg_available_extensions WHERE name = 'pg_hashids';


Usage
============
#### Encoding
Returns a hash using the default `alphabet` and empty `salt`.

	SELECT id_encode(1001); -- Result: jNl

Returns a hash using the default `alphabet` and supplied `salt`.

	SELECT id_encode(1234567, 'This is my salt'); -- Result: Pdzxp

Returns a hash using the default `alphabet`, `salt` and minimum hash length.
	
	SELECT id_encode(1234567, 'This is my salt', 10); -- Result: PlRPdzxpR7
	

Returns a hash using the supplied `alphabet`, `salt` and minimum hash length.
	
	SELECT id_encode(1234567, 'This is my salt', 10, 'abcdefghijABCDxFGHIJ1234567890'); -- Result: 3GJ956J9B9

#### Decoding
You can also decode previouslt generated hashes. Just use the same `salt`, otherwise you'll get wrong results.

	SELECT id_decode('PlRPdzxpR7', 'This is my salt', 10); -- Result: 1234567
	
Using a custom alphabet

	SELECT id_decode('3GJ956J9B9', 'This is my salt', 10, 'abcdefghijABCDxFGHIJ1234567890'); -- Result: 1234567

