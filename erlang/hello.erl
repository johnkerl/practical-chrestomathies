% erlc hello.erl
% erl -noshell -s hello -s init stop
-module(hello).
-export([start/0]).

start() ->
	io:fwrite("Hello in Erlang!\n").
