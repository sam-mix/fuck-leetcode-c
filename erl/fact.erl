-module(fact).

-export([fact/1, t/0, tt/1]).

fact(1) -> 1;
fact(N) -> fact(N - 1) * N.

t() -> fact(10).

tt(N) when N =< 0 ->
    ok;
tt(N) ->
    io:format("~p~n", [fact(N)]),
    tt(N - 1).
