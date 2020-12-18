//
//  MainAppView.swift
//  Ida
//
//  Created by Aditya Duri on 12/18/20.
//

import SwiftUI
import IdaParser

struct MainAppView: View {
    
    @Binding var voterName: String
    @Binding var voterMsg: String
    @Binding var parser: Parser?
    @Binding var parsedVotes: ParsedVotes?
    
    var body: some View {
        VStack(alignment: .leading) {
            HStack {
                VStack(alignment: .leading) {
                    TextField("Voter Name", text: $voterName)
                        .disableAutocorrection(true)
                    TextEditor(text: $voterMsg)
                        .disableAutocorrection(true)
                }
                .frame(minWidth: 0, maxWidth: .infinity)
                Divider()
                VStack(alignment: .leading) {
                    ParsedVotesView(parsedVotes: parsedVotes).padding(.horizontal)
                    if let parsedVotes = self.parsedVotes {
                        if parsedVotes.warningMsgs.count > 0 {
                            Divider()
                        }
                    }
                    VStack(alignment: .leading) {
                        if let warningMsgs = self.parsedVotes?.warningMsgs {
                            ForEach(warningMsgs, id: \.self) { warning in
                                Label(warning, systemImage: "exclamationmark.triangle.fill")
                                    .font(.caption)
                            }
                        }
                    }
                }
                .frame(minWidth: 0, maxWidth: .infinity)
            }
        }
    }
}

#if DEBUG
struct MainAppView_PreviewWrapper: View {
    @State var voterName: String = "Germany"
    @State var voterMsg: String = """
    12 :spain:
    10 :unitedkingdom:
    8 :thenetherlands:
    7 :israel:
    6 :portugal:
    5 :romania:
    4 :russia:
    3 :macedonia:
    2 :greece:
    1 :finland:
    """
    
    @State var parser: Parser? = nil
    @State var parsedVotes: ParsedVotes? = ParsedVotes(
        votes: [
            Vote(entry: Entry(country: Country(forum: "estonia", names: ["Estonia"]), artist: "Ines", song: "Once in a Lifetime"), points: 12),
            Vote(entry: Entry(country: Country(forum: "latvia", names: ["Latvia"]), artist: "Brainstorm", song: "My Star"), points: 10),
            Vote(entry: Entry(country: Country(forum: "russia", names: ["Russia"]), artist: "Alsou", song: "Solo"), points: 8)
        ],
        warningMsgs: [
            "Total number of points was not 58: 12"
        ]
    )
    
    var body: some View {
        MainAppView(voterName: $voterName, voterMsg: $voterMsg, parser: $parser, parsedVotes: $parsedVotes)
    }
}

struct MainAppView_Previews: PreviewProvider {
    static var previews: some View {
        MainAppView_PreviewWrapper()
    }
}
#endif
