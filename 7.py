lines = open(0).read().splitlines()
vs = dict()
elo = dict()
elo_default = 1200
for line in lines:
    a, b, p = line.split(',')
    ap, bp = [int(_) for _ in p.split('-')]
    left_wins = (ap > bp)
    if a not in vs:
        vs[a] = {}
        elo[a] = elo_default
    if b not in vs:
        vs[b] = {}
        elo[b] = elo_default
    vs[a][b] = left_wins
    vs[b][a] = not left_wins
    print(a, b, ap, bp, left_wins)
    # Elo:
    #   Ea = 1 / (1 + 10 ^ ((Rb - Ra) / 400))
    #   Ri' = Ri + 20 * (1 - Ei)
    def expected_win_rate(Ra, Rb):
        return 1 / (1 + 10 ** ((Rb - Ra) / 400))
    def update_score(E, R, has_won):
        return R + 20 * (int(has_won)- E)

    Ra = elo[a]
    Rb = elo[b]

    Ea = expected_win_rate(Ra, Rb)
    elo[a] = update_score(Ea, Ra, left_wins)

    Eb = expected_win_rate(Rb, Ra)
    elo[b] = update_score(Eb, Rb, not left_wins)

elo_sorted = dict(sorted(elo.items(), key=lambda _: _[1], reverse=True))
for k, v in elo_sorted.items():
    print(k, v)


