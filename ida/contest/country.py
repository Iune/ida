class Country:
    def __init__(self, forum_code, names):
        self.forum_code = forum_code
        self.names = names
        self.lower_names = [name.lower() for name in self.names]
        
        try:
            self.primary_name = self.names[0]
        except IndexError:
            raise ValueError("Country must have at least one name")

    def contains(self, name):
        if name.lower() in self.lower_names:
            return True
        return False
    