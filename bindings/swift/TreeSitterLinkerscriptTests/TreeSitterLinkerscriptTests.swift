import XCTest
import SwiftTreeSitter
import TreeSitterLinkerscript

final class TreeSitterLinkerscriptTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_linkerscript())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Linkerscript grammar")
    }
}
