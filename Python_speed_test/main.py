import pymongo, string , random
import time, threading

size=500000

def main():
    print(pymongo.__version__)
    myclient = pymongo.MongoClient("mongodb://root:example@localhost:37017/")
    mydb = myclient["mydb"]
    mycol = mydb["pythontest"]
    print("connection with DB ok")
    print("Generating Datasets")
    dataset1 = generateDataset(size)
    dataset2 = generateDataset(size)
    dataset3 = generateDataset(size)
    print("Datasets generated..")
    insertsingle(mycol,dataset1)
    insertsingleparall(mycol,dataset2)
    insertbulk(mycol,dataset3)
    findingtest(mycol)

def insertbulk(mycol, mylist):
    print("Staring inserting bulk...")
    start = time.time()
    x = mycol.insert_many(mylist)
    stop= time.time()
    print("finished insert bulk. -- Duration:", str(stop-start))


def insertsingle(mycol, mylist):
    print("Staring inserting single without multithread...")
    start = time.time()
    for i in mylist:
        x = mycol.insert_one(i)
    stop = time.time()
    print("finished insert single no parallelism -- Duration: ", str(stop - start))

def insertsingleparall(mycol, mylist):
    print("Staring inserting single with multithread...")
    threads=[]
    start = time.time()
    for i in mylist:
        x = threading.Thread(target=mycol.insert_one, args=(i,))
        threads.append(x)
        x.start()

    for index, thread in enumerate(threads):
        thread.join()
    stop = time.time()
    print("finished insert single with parallelism.  -- Duration:", str(stop - start))


def findingtest(mycol):
    print("Staring findingone 10 times...")
    filter = {"age": 50}
    start = time.time()
    for i in range(10):
        x = mycol.find_one(filter)
        print(x)
    stop = time.time()
    print("finished findingone(filter)x10.Duration: ", str(stop - start))

    print("Staring counting docs...")
    start = time.time()
    doc_count = mycol.count_documents(filter)
    print(doc_count)
    stop = time.time()
    print("finished Countdocument()..Duration: ", str(stop - start))

    print("Staring specific find...")
    start = time.time()
    myfound = mycol.find(filter)
    i=0
    for post in myfound:
        print(post)
        i+=1
        if i==10:
            break
    stop = time.time()
    print("finished Find(filter)..Duration: ", str(stop - start))



def generateDataset(ofsize):
    arra = []
    for i in range(ofsize):
        arra.append({"name": randomString(), "age":random.randint(1, 101), "city":randomString()})
    return arra

def randomString(stringLength=13):
    """Generate a random string of fixed length """
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(stringLength))



main()
