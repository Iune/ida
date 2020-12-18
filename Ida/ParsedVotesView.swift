//
//  ParsedVotesView.swift
//  Ida
//
//  Created by Aditya Duri on 12/18/20.
//

import SwiftUI
import IdaParser

struct ParsedVotesView: View {
    var parsedVotes: ParsedVotes?
    @State private var width: CGFloat? = nil
    
    var body: some View {
        ScrollView {
            LazyVStack(alignment: .leading) {
                ForEach(self.parsedVotes?.votes ?? [], id: \.self) { vote in
                    VoteRow(vote: vote, width: width)
                        .background(RoundedRectangle(cornerRadius: 20, style: .continuous).fill(Color.white))
                    Spacer()
                }
            }
        }
        .frame(minWidth: 0, maxWidth: .infinity, alignment: .topLeading)
        .onPreferenceChange(CenteringColumnPreferenceKey.self) { preferences in
            for p in preferences {
                let oldWidth = self.width ?? CGFloat.zero
                if p.width > oldWidth {
                    self.width = p.width * 1.5
                }
            }
        }
    }
}

struct VoteRow: View {
    var vote: Vote
    var width: CGFloat? = nil
    
    var body: some View {
        HStack {
            Text("\(vote.points)").font(.subheadline)
                .fontWeight(.black)
                .frame(width: width)
                .lineLimit(1)
                .background(CenteringView())
                .background(Capsule().fill(Color.black))
                .foregroundColor(.white)
            VStack(alignment: .leading) {
                Text(vote.entry.country.primaryName)
                    .bold()
                    .lineLimit(1)
                Text("\(vote.entry.artist) – \(vote.entry.song)")
                    .font(.caption)
                    .lineLimit(1)
                    .foregroundColor(.secondary)
            }
        }
        .frame(minWidth: 0, maxWidth: .infinity, alignment: .topLeading)
        .padding(.leading)
    }
}

/// This approach to centering items in a `VStack` was taken from https://stackoverflow.com/a/57677582
struct CenteringView: View {
    var body: some View {
        GeometryReader { geometry in
            Rectangle()
                .fill(Color.clear)
                .preference(
                    key: CenteringColumnPreferenceKey.self,
                    value: [CenteringColumnPreference(width: geometry.frame(in: CoordinateSpace.global).width)]
                )
        }
    }
}

struct CenteringColumnPreference: Equatable {
    let width: CGFloat
}

struct CenteringColumnPreferenceKey: PreferenceKey {
    typealias Value = [CenteringColumnPreference]
    
    static var defaultValue: [CenteringColumnPreference] = []
    
    static func reduce(value: inout [CenteringColumnPreference], nextValue: () -> [CenteringColumnPreference]) {
        value.append(contentsOf: nextValue())
    }
}

private var parsedVotes = ParsedVotes(
    votes: [
        Vote(entry: Entry(country: Country(forum: "estonia", names: ["Estonia"]), artist: "Ines", song: "Once in a Lifetime"), points: 12),
        Vote(entry: Entry(country: Country(forum: "latvia", names: ["Latvia"]), artist: "Brainstorm", song: "My Star"), points: 10),
        Vote(entry: Entry(country: Country(forum: "russia", names: ["Russia"]), artist: "Alsou", song: "Solo"), points: 8)
    ], warningMsgs: [])

struct ParsedVotesView_Previews: PreviewProvider {
    static var previews: some View {
        ParsedVotesView(parsedVotes: parsedVotes)
            .padding(.leading)
    }
}
