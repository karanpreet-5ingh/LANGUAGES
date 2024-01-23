'''Things to keep in mind 
1. self is just a pointer pointing to a class 
2. self is not a pre-define  keyword.
3. you can also use any thiong other than self 
4. __init__ is the inbuilt  function which is use formulate the data which is we are going to use inthe class
5. __str__ is a inbuilt function which is called when we use print(karan) the instance of class.


'''



class student:

    def __init__(self, name, rollno, joining_date, current_topic):
        self.name = name 
        self.rollno =  rollno
        self.joining_date =  joining_date
        self.current_topic = current_topic

    def name_parsing(self):
        if type(self.name) == list:
            for i in self.name:
                print("name od the student is : ", i)
        
        else:
            print("provided name is not in the form of list")

    def crt_topic(self):
        print("current topic is: ", self.current_topic)


    def str_rollno(self):


        try:
            if type(self.rollno) == str:
                print("nothing to do!!!")
            else:
                return str(self.rollno)
        except Exception as e:

            print("this is my error message",e)


    def duration(self, current_date):
        print("duration of the current topic is: ", current_date - self.joining_date)


    def __str__(self):
        return"this is the student class where we are trying to input their own data and they can try to fetch it later"





karan = student("karan",23,2023, "oop")


karan.duration(2024)

print(karan.name)