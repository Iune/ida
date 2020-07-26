import XCTest

import idaTests

var tests = [XCTestCaseEntry]()
tests += idaTests.allTests()
XCTMain(tests)
