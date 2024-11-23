from abc import ABC, abstractmethod
import copy
from enum import Enum, unique

# Prototype
class Prototype(ABC):
    _prototype_key = None
    _prototype_value = None

    def __init__(self, prototype_key: str = None, prototype_value: int = None, prototype = None):
        self._prototype_key = prototype_key
        self._prototype_value = prototype_value

    @abstractmethod
    def clone(self) -> 'Prototype':
        pass

    def print(self):
        print("Called from key: "+ self._prototype_key+", value: "+ str(self._prototype_value))

# Concrete Prototype
class ConcretePrototype1(Prototype):
    concrete_prototype_value_1 = None

    def __init__(self, prototype_key: str = None, prototype_value: int = None, concrete_prototype_1 = None):
        super(ConcretePrototype1, self).__init__(prototype_key,prototype_value,concrete_prototype_1)
        if isinstance(concrete_prototype_1, ConcretePrototype1) and concrete_prototype_1 is not None:
            self.concrete_prototype_value_1 = concrete_prototype_1.concrete_prototype_value_1

    def clone(self) -> 'Prototype':
        return copy.deepcopy(self)
        # return copy.copy(self)

    def __copy__(self):
            """
            Create a shallow copy. This method will be called whenever someone calls
            `copy.copy` with this object and the returned value is returned as the
            new shallow copy.
            """
            concrete_prototype_value_1 = copy.copy(self.concrete_prototype_value_1)
            new = self.__class__(concrete_prototype_value_1)
            new.__dict__.update(self.__dict__)
            return new

    def __deepcopy__(self, memo=None):
        """
        Create a deep copy. This method will be called whenever someone calls
        `copy.deepcopy` with this object and the returned value is returned as
        the new deep copy.

        What is the use of the argument `memo`? Memo is the dictionary that is
        used by the `deepcopy` library to prevent infinite recursive copies in
        instances of circular references. Pass it to all the `deepcopy` calls
        you make in the `__deepcopy__` implementation to prevent infinite
        recursions.
        """
        if memo is None:
            memo = {}
        concrete_prototype_value_1 = copy.deepcopy(self.concrete_prototype_value_1, memo)
        new = self.__class__(concrete_prototype_value_1)
        new.__dict__ = copy.deepcopy(self.__dict__, memo)
        return new
    
    def print(self):
        print("Called from key: "+ self._prototype_key+", value: "+ str(self._prototype_value)+", concrete_prototype_value_1: "+str(self.concrete_prototype_value_1))

class ConcretePrototype2(Prototype):
    concrete_prototype_value_2 = None

    def __init__(self, prototype_key: str = None, prototype_value: int = None, concrete_prototype_2 = None):
        super(ConcretePrototype2, self).__init__(prototype_key,prototype_value,concrete_prototype_2)
        if isinstance(concrete_prototype_2, ConcretePrototype2) and concrete_prototype_2 is not None:
            self.concrete_prototype_value_2 = concrete_prototype_2.concrete_prototype_value_2
    
    def clone(self) -> 'Prototype':
        return copy.deepcopy(self)
        # return copy.copy(self)

    def __copy__(self):
            """
            Create a shallow copy. This method will be called whenever someone calls
            `copy.copy` with this object and the returned value is returned as the
            new shallow copy.
            """
            concrete_prototype_value_2 = copy.copy(self.concrete_prototype_value_2)
            new = self.__class__(concrete_prototype_value_2)
            new.__dict__.update(self.__dict__)
            return new

    def __deepcopy__(self, memo=None):
        """
        Create a deep copy. This method will be called whenever someone calls
        `copy.deepcopy` with this object and the returned value is returned as
        the new deep copy.

        What is the use of the argument `memo`? Memo is the dictionary that is
        used by the `deepcopy` library to prevent infinite recursive copies in
        instances of circular references. Pass it to all the `deepcopy` calls
        you make in the `__deepcopy__` implementation to prevent infinite
        recursions.
        """
        if memo is None:
            memo = {}
        concrete_prototype_value_2 = copy.deepcopy(self.concrete_prototype_value_2, memo)
        new = self.__class__(concrete_prototype_value_2)
        new.__dict__ = copy.deepcopy(self.__dict__, memo)
        return new
    
    def print(self):
        print("Called from key: "+ self._prototype_key+", value: "+ str(self._prototype_value)+", concrete_prototype_value_2: "+str(self.concrete_prototype_value_2))

# Prototype Registry Keys
@unique
class PrototypeKey(Enum):
   PROTOTYPE_1 = "PROTOTYPE_1"
   PROTOTYPE_2 = "PROTOTYPE_2"


# Prototype manager
class PrototypeManager:
    _prototypes = {}

    def __init__(self):
        concrete_prototype = ConcretePrototype1(PrototypeKey.PROTOTYPE_1.name, 1)
        concrete_prototype.concrete_prototype_value_1 = 11
        self._prototypes[PrototypeKey.PROTOTYPE_1.name] = concrete_prototype
    
    def create_prototype(self, prototype_key: str) -> Prototype:
        print("Creating prototype: "+prototype_key)
        return self._prototypes[prototype_key].clone()
    
    def list_prototypes(self):
        print("Listing prototypes")
        for k in self._prototypes.keys():
            print(k) 

    def put_prototype(self, prototype_key: str, prototype: Prototype):
        print("Putting prototype <" + prototype_key+ ">")
        self._prototypes[prototype_key] = prototype

    def compare_with_prototype(self, prototype_key: str, prototype: Prototype):
        print("Comparing with Prototype:")
        if id(self._prototypes[prototype_key]) == id(prototype):
            print("- Same objects")
        else:
            print("- Different objects")
        if type(self._prototypes[prototype_key]).__name__  == type(prototype).__name__ and self._prototypes[prototype_key]._prototype_key == prototype._prototype_key:
            print("- Identical")
        else:
            print("- Not Identical")

    def compare_with_clone(self, prototype: Prototype, prototype2: Prototype):
        print("Comparing with Prototype:")
        if id(prototype) == id(prototype2):
            print("- Same objects")
        else:
            print("- Different objects")
        if type(prototype).__name__  == type(prototype2).__name__ and prototype._prototype_key == prototype2._prototype_key:
            print("- Identical")
        else:
            print("- Not Identical")

def client_code():
    prototype_manager = PrototypeManager()
    prototype_manager.list_prototypes()
    concrete_prototype2 = ConcretePrototype2(PrototypeKey.PROTOTYPE_2.name, 2)
    concrete_prototype2.concrete_prototype_value_2 = 22
    prototype_manager.put_prototype(PrototypeKey.PROTOTYPE_2.name,concrete_prototype2)
    prototype_manager.list_prototypes()

    print("======PROTOTYPE_1========")
    prototype1 = prototype_manager.create_prototype(PrototypeKey.PROTOTYPE_1.name)
    prototype1.print()
    prototype_manager.compare_with_prototype(PrototypeKey.PROTOTYPE_1.name,prototype1)

    print("")
    prototype2 = prototype_manager.create_prototype(PrototypeKey.PROTOTYPE_1.name)
    prototype2.print()
    prototype_manager.compare_with_prototype(PrototypeKey.PROTOTYPE_1.name,prototype2)

    print("")
    prototype_manager.compare_with_clone(prototype1, prototype2)

    print("")
    print("======PROTOTYPE_2========")
    prototype3 = prototype_manager.create_prototype(PrototypeKey.PROTOTYPE_2.name)
    prototype3.print()
    prototype_manager.compare_with_prototype(PrototypeKey.PROTOTYPE_2.name,prototype3)

    print("")
    prototype4 = prototype_manager.create_prototype(PrototypeKey.PROTOTYPE_2.name)
    prototype4.print()
    prototype_manager.compare_with_prototype(PrototypeKey.PROTOTYPE_2.name,prototype4)

    print("")
    prototype_manager.compare_with_clone(prototype3, prototype4)

    print("======DIFF========")
    prototype_manager.compare_with_clone(prototype1, prototype3)
    prototype_manager.compare_with_clone(prototype2, prototype4)
    

if __name__ == "__main__":
    client_code()