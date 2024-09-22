class DIE:
    def __init__(self, initial_state):
        self.faces = initial_state
    def Show(self, idx: int = None):
        print(idx, self.faces) if idx is not None else print(self.faces)
    def Faces(self):
        return self.faces
    def Roll(self, op):
        if op == 'U': self.Roll_Up()
        elif op == 'D': self.Roll_Down()
        elif op == 'L': self.Roll_Left()
        elif op == 'R': self.Roll_Right()
    def Roll_Up(self):
        up = self.faces['front']
        front = self.faces['down']
        self.faces['up'] = up
        self.faces['down'] = 7 - up
        self.faces['front'] = front
        self.faces['back'] = 7 - front
    def Roll_Down(self):
        up = self.faces['back']
        front  = self.faces['up']
        self.faces['up'] = up
        self.faces['down'] = 7 - up
        self.faces['front'] = front
        self.faces['back'] = 7 - front
    def Roll_Left(self):
        left = self.faces['front']
        front = self.faces['right']
        self.faces['left'] = left
        self.faces['right'] = 7 - left
        self.faces['front'] = front
        self.faces['back'] = 7 - front
    def Roll_Right(self):
        left = self.faces['back']
        front = self.faces['left']
        self.faces['left'] = left
        self.faces['right'] = 7 - left
        self.faces['front'] = front
        self.faces['back'] = 7 - front

db = DIE({'up': 2, 'front': 1, 'left': 3,
    'down': 7-2, 'back': 7-1, 'right': 7-3,
})

da = DIE({'up': 3, 'front': 1, 'left': 2,
    'down': 7-3, 'back': 7-1, 'right': 7-2,
})

da.Show()
db.Show()

ops = open(0).read().strip()

# TEST
#ops = 'LRDLU'

res = 0
matches = []
for i, op in enumerate(ops):
    da.Roll(op)
    db.Roll(op)
    da.Show(i)
    db.Show(i)
    if da.Faces()['front'] == db.Faces()['front']:
        matches.append(i)
        res += i

da.Show()
db.Show()
print('matches/', matches)
print('res/', res)

